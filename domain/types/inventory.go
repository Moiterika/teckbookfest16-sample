package types

import decimal "github.com/shopspring/decimal"

// Inventory 型は在庫、あるいは、在庫増減を表す。
type Inventory struct {
	val  decimal.Decimal // マイナスを許可する
	unit Unit
}

// Val は按分用のメソッド
func (i Inventory) Val() decimal.Decimal {
	return i.val
}

// Unit は按分用のメソッド
func (i Inventory) Unit() Unit {
	return i.unit
}

func NewInventory(数量 decimal.Decimal, 数量単位 Unit) Inventory {
	return Inventory{
		val:  数量,
		unit: 数量単位,
	}
}
