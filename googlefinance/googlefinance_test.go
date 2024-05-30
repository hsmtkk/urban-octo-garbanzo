package googlefinance_test

import (
	"os"
	"testing"

	"github.com/hsmtkk/urban-octo-garbanzo/googlefinance"
	"github.com/stretchr/testify/assert"
)

func TestGetPrice(t *testing.T) {
	apiKey := os.Getenv("SERP_API_KEY")
	assert.NotEmpty(t, apiKey)
	price, err := googlefinance.New(apiKey).GetPrice("CBOT:XWN24")
	assert.Nil(t, err)
	assert.Positive(t, price)
}
