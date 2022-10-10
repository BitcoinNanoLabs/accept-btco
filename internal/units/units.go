package units

import (
	"github.com/shopspring/decimal"
)

const exp = 30

func NanoToRaw(btco decimal.Decimal) decimal.Decimal {
	return btco.Shift(exp)
}

func RawToNano(raw decimal.Decimal) decimal.Decimal {
	return raw.Shift(-exp)
}
