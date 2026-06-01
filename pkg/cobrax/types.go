package cobrax

import (
	"github.com/spf13/cobra"
)

type Command struct {
	*cobra.Command
	Scheduler string
}

func NewCommand(command *cobra.Command, scheduler string) *Command {
	return &Command{Command: command, Scheduler: scheduler}
}
