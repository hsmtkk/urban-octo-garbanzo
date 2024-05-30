package symbol_test

import (
	"testing"
	"time"

	"github.com/hsmtkk/wheat-future-sim/symbol"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	begin, err := time.Parse("2006-01-02", "2024-05-30")
	assert.Nil(t, err)
	end := time.Now().AddDate(1, 0, 0)
	got := symbol.Generate(begin, end)
	want := []string{"XWN24", "XWU24", "XWZ24", "XWH25", "XWK25"}
	assert.Equal(t, len(want), len(got))
	for i := 0; i < len(want); i++ {
		assert.Equal(t, want[i], got[i].Symbol)
	}
}
