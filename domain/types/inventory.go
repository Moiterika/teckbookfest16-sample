package types

import (
	"encoding/json"

	decimal "github.com/shopspring/decimal"
)

// Inventory 型は在庫、あるいは、在庫増減を表す。
type Inventory struct {
	val  decimal.Decimal // マイナスを許可する
	unit Code単位
}

func NewInventory(数量 decimal.Decimal, 数量単位 Code単位) Inventory {
	return Inventory{
		val:  数量,
		unit: 数量単位,
	}
}

// Val は按分用のメソッド
func (i Inventory) Val() decimal.Decimal {
	return i.val
}

// Unit は按分用のメソッド
func (i Inventory) Unit() Code単位 {
	return i.unit
}

func (i Inventory) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Val  decimal.Decimal `json:"val"`
		Unit Code単位          `json:"unit"`
	}{
		Val:  i.val,
		Unit: i.unit,
	})
	return v, err
}

func (i *Inventory) UnmarshalJSON(byte []byte) error {
	var s struct {
		Val  decimal.Decimal `json:"val"`
		Unit Code単位          `json:"unit"`
	}

	err := json.Unmarshal(byte, &s)
	if err != nil {
		return err
	}

	*i = NewInventory(s.Val, s.Unit)

	return nil
}
