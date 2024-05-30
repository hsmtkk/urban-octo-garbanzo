package rollover

import (
	"time"

	"github.com/hsmtkk/urban-octo-garbanzo/symbol"
)

/*
M: 現在～第1限月の最終取引日の週数
N: 第1限月の最終取引日～第2限月の最終取引日の週数

- 売り
  - 第1限月M枚、第2限月(N-M)枚
- 買い
  - 第3限月M枚、第4限月(N-M)枚
*/

type SymbolQuantity struct {
	Symbol   string
	Expire   time.Time
	Quantity int
}

type CalculateResult struct {
	Sell1 SymbolQuantity
	Sell2 SymbolQuantity
	Buy3  SymbolQuantity
	Buy4  SymbolQuantity
}

func Calculate(begin time.Time, end time.Time) CalculateResult {
	symbolMonths := symbol.Generate(begin, end)
	symbolMonth1 := symbolMonths[0]
	symbolMonth2 := symbolMonths[1]
	symbolMonth3 := symbolMonths[2]
	symbolMonth4 := symbolMonths[3]
	diffBegin1 := symbolMonth1.Expire.Sub(begin) // 始点～第1限月
	m := int(diffBegin1.Hours()) / 24 / 7
	diff12 := symbolMonth2.Expire.Sub(symbolMonth1.Expire) // 第1限月～第2限月
	n := int(diff12.Hours()) / 24 / 7

	return CalculateResult{
		Sell1: SymbolQuantity{symbolMonth1.Symbol, symbolMonth1.Expire, m},
		Sell2: SymbolQuantity{symbolMonth2.Symbol, symbolMonth2.Expire, n - m},
		Buy3:  SymbolQuantity{symbolMonth3.Symbol, symbolMonth3.Expire, m},
		Buy4:  SymbolQuantity{symbolMonth4.Symbol, symbolMonth4.Expire, n - m},
	}
}
