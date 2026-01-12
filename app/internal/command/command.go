package command

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"go-zero-box/app/internal/command/demo"
	"go-zero-box/pkg/asynqx"
)

var Provider = wire.NewSet(
	NewCommand,
	demo.NewHello,
)

type Command struct {
	Hello *demo.Hello
}

func NewCommand(hello *demo.Hello) *Command {
	return &Command{Hello: hello}
}

// RegisterHandlerScript 注册脚本任务
func RegisterHandlerScript(s *Command) []*cobra.Command {
	return []*cobra.Command{
		s.Hello.Sync(),
	}
}

// RegisterHandlerScheduler 注册计划任务
func RegisterHandlerScheduler(s *Command) []*asynqx.Handler {
	return []*asynqx.Handler{
		s.Hello.Async(),
	}
}
