// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	decimal "github.com/shopspring/decimal"
	objects "techbookfest16-sample/domain/objects"
	"time"
)

type Dto受払 struct {
	FldNo     Id
	Fld登録日時   time.Time
	Fld計上月    time.Time
	Fld受払区分   objects.Enum受払区分
	Fld赤伝フラグ  bool
	Fld品目ID   Id
	Fld基準数量   decimal.Decimal
	Fld基準単位ID Id

	rowState DataRowState
	Ub       *ub受払 `json:"-"`
}

func (d Dto受払) TableName() string {
	return "受払"
}
func (d Dto受払) RowState() DataRowState {
	return d.rowState
}

// Import はDto受払型に主キー以外を上書きする。
func (d *Dto受払) Import(登録日時 time.Time, 計上月 time.Time, 受払区分 objects.Enum受払区分, 赤伝フラグ bool, 品目ID Id, 基準数量 decimal.Decimal, 基準単位ID Id) {
	// 項目がすべて一致していたら、何もしない
	if d.Fld登録日時 == 登録日時 && d.Fld計上月 == 計上月 && d.Fld受払区分 == 受払区分 && d.Fld赤伝フラグ == 赤伝フラグ && d.Fld品目ID == 品目ID && d.Fld基準数量 == 基準数量 && d.Fld基準単位ID == 基準単位ID {
		return
	}
	if d.Fld登録日時 != 登録日時 {
		d.Fld登録日時 = 登録日時
		d.Ub.Set(Tbl受払().Fld登録日時(), 登録日時)
	}
	if d.Fld計上月 != 計上月 {
		d.Fld計上月 = 計上月
		d.Ub.Set(Tbl受払().Fld計上月(), 計上月)
	}
	if d.Fld受払区分 != 受払区分 {
		d.Fld受払区分 = 受払区分
		d.Ub.Set(Tbl受払().Fld受払区分(), 受払区分)
	}
	if d.Fld赤伝フラグ != 赤伝フラグ {
		d.Fld赤伝フラグ = 赤伝フラグ
		d.Ub.Set(Tbl受払().Fld赤伝フラグ(), 赤伝フラグ)
	}
	if d.Fld品目ID != 品目ID {
		d.Fld品目ID = 品目ID
		d.Ub.Set(Tbl受払().Fld品目ID(), 品目ID)
	}
	if d.Fld基準数量 != 基準数量 {
		d.Fld基準数量 = 基準数量
		d.Ub.Set(Tbl受払().Fld基準数量(), 基準数量)
	}
	if d.Fld基準単位ID != 基準単位ID {
		d.Fld基準単位ID = 基準単位ID
		d.Ub.Set(Tbl受払().Fld基準単位ID(), 基準単位ID)
	}

}
