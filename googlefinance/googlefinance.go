package googlefinance

import (
	"os"

	"github.com/hsmtkk/urban-octo-garbanzo/myhttp"
)

const URL = "https://serpapi.com/search?engine=google_finance"

type PriceGetter interface {
	GetPrice(symbol string) (float64, error)
}

func New(apiKey string) PriceGetter {
	return &priceGetter{apiKey: apiKey}
}

type priceGetter struct {
	apiKey string
}

func (p *priceGetter) GetPrice(symbol string) (float64, error) {
	query := map[string]string{
		"api_key": p.apiKey,
		"engine":  "google_finance",
		"q":       symbol,
	}
	content, err := myhttp.Get(URL, query)
	if err != nil {
		return 0, err
	}
	os.WriteFile("./sample.json", content, 0644)
	return 0, nil
}
