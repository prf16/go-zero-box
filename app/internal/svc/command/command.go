package command

import (
	"go-zero-box/app/internal/svc/command/hello"
	"go-zero-box/pkg/cobrax"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewCommand,
	hello.NewWorld,
)

type Command struct {
	World *hello.World
}

func NewCommand(hello *hello.World) *Command {
	return &Command{World: hello}
}

// Register 注册任务
func (c *Command) Register() []*cobrax.Command {
	return []*cobrax.Command{
		{Command: c.World.Handle(), Scheduler: "* * * * *"},
	}
}
