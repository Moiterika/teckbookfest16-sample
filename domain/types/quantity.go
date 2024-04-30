package types

import (
	"encoding/json"

	decimal "github.com/shopspring/decimal"
	"golang.org/x/xerrors"
)

// Quantity 型は数量を表す。数量は常に0以上の値である。
type Quantity struct {
	val  decimal.Decimal
	unit Code単位
}

// NewQuantity は数量（Quantity型）を生成する。マイナスの場合、エラーとなる。
func NewQuantity(数量 decimal.Decimal, 数量単位 Code単位) (q Quantity, err error) {
	// 数量 < 0
	if 数量.LessThan(decimal.Zero) {
		err = xerrors.Errorf("不正な値。数量=%s", 数量.String())
		return
	}
	q = Quantity{
		val:  数量,
		unit: 数量単位,
	}
	return
}

// Val は按分用のメソッド
func (q Quantity) Val() decimal.Decimal {
	return q.val
}

// Unit は按分用のメソッド
func (q Quantity) Unit() Code単位 {
	return q.unit
}

func (q Quantity) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Val  decimal.Decimal `json:"val"`
		Unit Code単位          `json:"unit"`
	}{
		Val:  q.val,
		Unit: q.unit,
	})
	return v, err
}

func (q *Quantity) UnmarshalJSON(byte []byte) error {
	var s struct {
		Val  decimal.Decimal `json:"val"`
		Unit Code単位          `json:"unit"`
	}

	err := json.Unmarshal(byte, &s)
	if err != nil {
		return err
	}

	*q, err = NewQuantity(s.Val, s.Unit)
	if err != nil {
		return err
	}

	return nil
}

// IsZero は数値が0 ⇒ true、0以外 ⇒ falseを返す。
func (q Quantity) IsZero() bool {
	return q.val.Equal(decimal.Zero)
}
