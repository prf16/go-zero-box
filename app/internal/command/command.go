package command

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"go-zero-box/app/internal/command/demo"
	"go-zero-box/pkg/asynqx"
)

var Provider = wire.NewSet(
	NewCommand,
	demo.NewPersonProcess,
)

type Command struct {
	PersonProcess *demo.PersonProcess
}

func NewCommand(personProcess *demo.PersonProcess) *Command {
	return &Command{PersonProcess: personProcess}
}

func CommandHandler(s *Command) []*cobra.Command {
	return []*cobra.Command{
		s.PersonProcess.Sync(),
	}
}

func SchedulerHandler(s *Command) []*asynqx.Handler {
	return []*asynqx.Handler{
		s.PersonProcess.Async(),
	}
}
