package updateprice

import (
	"fmt"
	"time"

	"github.com/hsmtkk/wheat-future-sim/symbol"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "updateprice",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	symbols := symbol.Generate(time.Now(), time.Now().AddDate(1, 0, 0))
	fmt.Println(symbols)
}
