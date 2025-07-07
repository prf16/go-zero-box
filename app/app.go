package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-box/app/internal/command"
	"go-zero-box/app/internal/config"
	"go-zero-box/app/internal/handler"
	"go-zero-box/app/internal/queue"
	"go-zero-box/pkg/asynqx"
	"log"
)

var (
	configFile = flag.String("conf", "app/etc/app.yaml", "the config file")

	rootCmd = &cobra.Command{
		Use:                "app",
		DisableFlagParsing: true,
		Hidden:             true,
	}
)

func main() {
	flag.Parse()
	var c *config.Config
	conf.MustLoad(*configFile, &c)
	logc.MustSetup(c.Server.Log)
	app := initApp(c)
	rootCmd.AddCommand(serverApi(app), serverQueue(app), serverScheduler(app), serverAll(app))
	rootCmd.AddCommand(command.ScriptHandler(app.command)...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("execute core service failed, %s\n", err.Error())
	}
}

func serverApi(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:api",
		Short: "启动api服务",
		Run: func(cmd *cobra.Command, args []string) {
			server := rest.MustNewServer(app.config.Server)
			defer server.Stop()
			handler.RegisterHandlers(server, app.svcCtx)
			//httpx.SetValidator(tools.NewValidator())
			log.Printf("[Server: api] start success at %s:%d...\n", app.config.Server.Host, app.config.Server.Port)
			server.Start()
		},
	}
}

func serverQueue(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:queue",
		Short: "启动队列服务",
		Run: func(cmd *cobra.Command, args []string) {
			serviceGroup := service.NewServiceGroup()
			defer serviceGroup.Stop()

			handlers := queue.Handler(app.queue)
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(app.config.Asynqx, v))
			}
			serviceGroup.Start()
			select {}
		},
	}
}

func serverScheduler(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:scheduler",
		Short: "启动计划任务服务",
		Run: func(cmd *cobra.Command, args []string) {
			serviceGroup := service.NewServiceGroup()
			defer serviceGroup.Stop()

			handlers := command.SchedulerHandler(app.command)
			serviceGroup.Add(asynqx.NewScheduler(app.config.Asynqx, handlers))
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(app.config.Asynqx, v))
			}

			serviceGroup.Start()
		},
	}
}

func serverAll(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:all",
		Short: "单体式服务，包含api、队列、计划任务",
		Run: func(cmd *cobra.Command, args []string) {
			go serverApi(app).Run(serverApi(app), []string{})
			go serverQueue(app).Run(serverQueue(app), []string{})
			go serverScheduler(app).Run(serverScheduler(app), []string{})
			select {}
		},
	}
}
