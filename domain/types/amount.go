package types

import (
	"encoding/json"
	"fmt"

	decimal "github.com/shopspring/decimal"
)

// Amount 型は金額を表す。
type Amount struct {
	val decimal.Decimal
	cur CurrencyUnit
}

// NewAmount は金額（Amount型）を生成する。金額は整数になるように四捨五入される。
func NewAmount(金額 decimal.Decimal, 通貨コード CurrencyUnit) Amount {
	return Amount{
		val: 金額.Round(0),
		cur: 通貨コード,
	}
}

// Val は按分用のメソッド
func (a Amount) Val() decimal.Decimal {
	return a.val
}

// Unit は按分用のメソッド
func (a Amount) Unit() CurrencyUnit {
	return a.cur
}

func (a Amount) String() string {
	return fmt.Sprintf("金額=%s、単位=%s", a.val.String(), a.cur.String())
}

func (a Amount) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Amt decimal.Decimal `json:"amt"`
		Cur CurrencyUnit    `json:"cur"`
	}{
		Amt: a.val,
		Cur: a.cur,
	})
	return v, err
}

func (a *Amount) UnmarshalJSON(byte []byte) error {
	var s struct {
		Amt decimal.Decimal `json:"amt"`
		Cur CurrencyUnit    `json:"cur"`
	}

	err := json.Unmarshal(byte, &s)
	if err != nil {
		return err
	}

	*a = NewAmount(s.Amt, s.Cur)

	return nil
}
