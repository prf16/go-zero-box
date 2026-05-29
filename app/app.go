package app

import (
	"context"
	"flag"
	"fmt"
	"go-zero-box/app/internal/config"
	"go-zero-box/app/internal/handler"
	"go-zero-box/app/internal/svc/queue"
	"go-zero-box/pkg/asynqx"
	"log"

	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile = flag.String("conf", "app/etc/app.yaml", "the config file")

	rootCmd = &cobra.Command{
		Use:                "app",
		DisableFlagParsing: true,
		Hidden:             true,
	}
)

func Start() {
	flag.Parse()
	var c *config.Config
	err := conf.Load(*configFile, &c)
	if err != nil {
		err = conf.Load("etc/app.yaml", &c)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
	}
	logc.MustSetup(c.Server.Log)

	app := initApp(c)

	rootCmd.AddCommand(
		serverApi(app),
		serverQueue(app),
		serverScheduler(app),
		serverAll(app),
	)

	for _, v := range app.svcCtx.Command.Register() {
		rootCmd.AddCommand(v.Command)
	}

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
			log.Printf("[server:api] start success at %s:%d...\n", app.config.Server.Host, app.config.Server.Port)
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

			handlers := queue.RegisterHandlerQueue(app.svcCtx.Queue)
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(app.config.Redis, v))
			}
			serviceGroup.Start()
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

			var handlers []*asynqx.Handler
			for _, v := range app.svcCtx.Command.Register() {
				if v.Scheduler == "" {
					continue
				}

				handlers = append(handlers, &asynqx.Handler{
					Type:      v.Command.Use,
					Scheduler: v.Scheduler,
					Async: func(ctx context.Context, task *asynq.Task) error {
						v.Command.Run(v.Command, nil)
						return nil
					},
				})
			}

			serviceGroup.Add(asynqx.NewScheduler(app.config.Redis, handlers))
			for _, v := range handlers {
				serviceGroup.Add(asynqx.NewQueue(app.config.Redis, v))
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
