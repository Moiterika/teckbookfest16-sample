package types

import (
	decimal "github.com/shopspring/decimal"
	"golang.org/x/xerrors"
)

// Quantity 型は数量を表す。数量は常に0以上の値である。
type Quantity struct {
	val  decimal.Decimal
	unit Code単位
}

// Val は按分用のメソッド
func (q Quantity) Val() decimal.Decimal {
	return q.val
}

// Unit は按分用のメソッド
func (q Quantity) Unit() Code単位 {
	return q.unit
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

// IsZero は数値が0 ⇒ true、0以外 ⇒ falseを返す。
func (q Quantity) IsZero() bool {
	return q.val.Equal(decimal.Zero)
}
