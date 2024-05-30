package rollover

import (
	"fmt"
	"os"
	"time"

	"github.com/hsmtkk/wheat-future-sim/rollover"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "rollover",
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	begin := time.Now()
	end := begin.AddDate(1, 0, 0)
	calcResult := rollover.Calculate(begin, end)

	fmt.Println("売り")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Symbol", "Expire", "Quantity"})
	addRow(t, calcResult.Sell1)
	addRow(t, calcResult.Sell2)
	t.Render()

	fmt.Println("買い")
	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Symbol", "Expire", "Quantity"})
	addRow(t, calcResult.Buy3)
	addRow(t, calcResult.Buy4)
	t.Render()
}

func addRow(t table.Writer, sq rollover.SymbolQuantity) {
	t.AppendRow([]interface{}{sq.Symbol, sq.Expire.Format("2006-01-02"), sq.Quantity})
}
