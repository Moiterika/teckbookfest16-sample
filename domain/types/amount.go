package types

import (
	"fmt"

	decimal "github.com/shopspring/decimal"
)

// Amount 型は金額を表す。
type Amount struct {
	val  decimal.Decimal
	unit CurrencyUnit
}

func (a Amount) Val() decimal.Decimal {
	return a.val
}

func (a Amount) Unit() CurrencyUnit {
	return a.unit
}

func (a Amount) String() string {
	return fmt.Sprintf("金額=%s、単位=%s", a.val.String(), a.unit.String())
}

// NewAmount は金額（Amount型）を生成する。金額は整数になるように四捨五入される。
func NewAmount(金額 decimal.Decimal, 通貨コード CurrencyUnit) Amount {
	return Amount{
		val:  金額.Round(0),
		unit: 通貨コード,
	}
}
