package types

import (
	"encoding/json"

	decimal "github.com/shopspring/decimal"
)

// Price 型は単価を表す。
type Price struct {
	amt     decimal.Decimal
	cur     CurrencyUnit
	perUnit Code単位
}

func NewPrice(amt decimal.Decimal, cur CurrencyUnit, perUnit Code単位) (p Price) {
	p.amt = amt.Round(6)
	p.cur = cur
	p.perUnit = perUnit
	return
}

func (p Price) Amt() decimal.Decimal {
	return p.amt
}
func (p Price) Cur() CurrencyUnit {
	return p.cur
}
func (p Price) PerUnit() Code単位 {
	return p.perUnit
}

func (p Price) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Amt  decimal.Decimal `json:"amt"`
		Cur  CurrencyUnit    `json:"cur"`
		Unit Code単位          `json:"per_unit"`
	}{
		Amt:  p.amt,
		Cur:  p.cur,
		Unit: p.perUnit,
	})
	return v, err
}
