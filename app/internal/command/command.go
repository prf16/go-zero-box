package command

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"go-zero-box/app/internal/command/demo"
	"go-zero-box/pkg/asynqx"
)

var Provider = wire.NewSet(
	NewCommand,
	demo.NewPerson,
)

type Command struct {
	Person *demo.Person
}

func NewCommand(person *demo.Person) *Command {
	return &Command{Person: person}
}

// ScriptHandler 注册脚本任务
func ScriptHandler(s *Command) []*cobra.Command {
	return []*cobra.Command{
		s.Person.Sync(),
	}
}

// SchedulerHandler 注册计划任务
func SchedulerHandler(s *Command) []*asynqx.Handler {
	return []*asynqx.Handler{
		s.Person.Async(),
	}
}
