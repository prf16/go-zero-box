package cobrax

import (
	"github.com/spf13/cobra"
)

type Command struct {
	Command   *cobra.Command
	Scheduler string
}
