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
	logc.MustSetup(c.Api.Log)
	config.Init(c)
	app := initApp()
	rootCmd.AddCommand(apiCommand(app), schedulerCommand(app), queueCommand(app))
	rootCmd.AddCommand(command.CommandHandler(app.command)...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("execute core service failed, %s\n", err.Error())
	}
}

func apiCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:api",
		Short: "启动api服务",
		Run: func(cmd *cobra.Command, args []string) {
			server := rest.MustNewServer(config.GetConfig().Api.RestConf)
			defer server.Stop()

			handler.RegisterHandlers(server, app.svcCtx)
			log.Printf("[Server: api] start success at %s:%d...\n", config.GetConfig().Api.Host, config.GetConfig().Api.Port)
			server.Start()
		},
	}
}

func queueCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:queue",
		Short: "启动队列服务",
		Run: func(cmd *cobra.Command, args []string) {
			serviceGroup := service.NewServiceGroup()
			defer serviceGroup.Stop()

			handlers := queue.Handler(app.queue)
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(config.GetConfig().Queue, v))
			}
			serviceGroup.Start()
			select {}
		},
	}
}

func schedulerCommand(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "server:scheduler",
		Short: "启动计划任务服务",
		Run: func(cmd *cobra.Command, args []string) {
			serviceGroup := service.NewServiceGroup()
			defer serviceGroup.Stop()

			handlers := command.SchedulerHandler(app.command)
			serviceGroup.Add(asynqx.NewScheduler(config.GetConfig().Scheduler, handlers))
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(config.GetConfig().Scheduler, v))
			}

			serviceGroup.Start()
		},
	}
}
