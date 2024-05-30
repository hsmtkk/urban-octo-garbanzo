package command

import (
	"github.com/hsmtkk/urban-octo-garbanzo/command/rollover"
	"github.com/hsmtkk/urban-octo-garbanzo/command/updateprice"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{}

func init() {
	Command.AddCommand(updateprice.Command)
	Command.AddCommand(rollover.Command)
}
