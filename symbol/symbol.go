package symbol

import "time"

type SymbolExpire struct {
	Symbol string
	Expire time.Time
}

type monthDay struct {
	Month int
	Day   int
}

// 指定期間内のシンボルを生成する
/*
- 2/28 H
- 4/30 K
- 6/30 N
- 8/31 U
- 11/30 Z
*/
func Generate(begin time.Time, end time.Time) []SymbolExpire {
	results := []SymbolExpire{}
	dateSymbolMap := map[monthDay]string{
		{2, 28}:  "H",
		{4, 30}:  "K",
		{6, 30}:  "N",
		{8, 31}:  "U",
		{11, 30}: "Z",
	}
	date := begin
	for {
		monthSymbol, contains := dateSymbolMap[monthDay{int(date.Month()), date.Day()}]
		if contains {
			symbol := "XW" + monthSymbol + date.Format("06")
			results = append(results, SymbolExpire{
				Symbol: symbol,
				Expire: date,
			})
		}
		date = date.AddDate(0, 0, 1)
		if date.After(end) {
			break
		}
	}
	return results
}
