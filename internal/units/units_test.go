package units

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const (
	rawValue  = "12340000000000000000000000000000"
	btcoValue = "12.34"
)

func TestNanoToRaw(t *testing.T) {
	i := NanoToRaw(decimal.RequireFromString(btcoValue))
	assert.Equal(t, rawValue, i.String())
}

func TestRawToNano(t *testing.T) {
	i := decimal.RequireFromString(rawValue)
	assert.Equal(t, btcoValue, RawToNano(i).String())
}
