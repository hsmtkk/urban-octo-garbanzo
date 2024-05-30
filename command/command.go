package command

import (
	"github.com/hsmtkk/wheat-future-sim/command/rollover"
	"github.com/hsmtkk/wheat-future-sim/command/updateprice"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{}

func init() {
	Command.AddCommand(updateprice.Command)
	Command.AddCommand(rollover.Command)
}
