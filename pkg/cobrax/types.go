package cobrax

import (
	"github.com/spf13/cobra"
)

type Command struct {
	*cobra.Command
	Scheduler string
}
