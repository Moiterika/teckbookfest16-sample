// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import decimal "github.com/shopspring/decimal"

type Dto受払製造実績 struct {
	FldNo     Id
	Fld製造数量   decimal.Decimal
	Fld製造単位ID Id
	Fld製造指図ID Id

	rowState DataRowState
	Ub       *ub受払製造実績 `json:"-"`
}

func (d Dto受払製造実績) TableName() string {
	return "受払_製造実績"
}
func (d Dto受払製造実績) RowState() DataRowState {
	return d.rowState
}

// Import はDto受払製造実績型に主キー以外を上書きする。
func (d *Dto受払製造実績) Import(製造数量 decimal.Decimal, 製造単位ID Id, 製造指図ID Id) {
	// 項目がすべて一致していたら、何もしない
	if d.Fld製造数量 == 製造数量 && d.Fld製造単位ID == 製造単位ID && d.Fld製造指図ID == 製造指図ID {
		return
	}
	if d.Fld製造数量 != 製造数量 {
		d.Fld製造数量 = 製造数量
		d.Ub.Set(Tbl受払製造実績().Fld製造数量(), 製造数量)
	}
	if d.Fld製造単位ID != 製造単位ID {
		d.Fld製造単位ID = 製造単位ID
		d.Ub.Set(Tbl受払製造実績().Fld製造単位ID(), 製造単位ID)
	}
	if d.Fld製造指図ID != 製造指図ID {
		d.Fld製造指図ID = 製造指図ID
		d.Ub.Set(Tbl受払製造実績().Fld製造指図ID(), 製造指図ID)
	}

}
