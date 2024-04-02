// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import objects "techbookfest16-sample/domain/objects"

type DtoEnum受払区分 struct {
	Fld受払区分  objects.Enum受払区分
	Fld名称    string
	Fld受入フラグ bool
	Fld符号    int

	rowState DataRowState
	Ub       *ubEnum受払区分 `json:"-"`
}

func (d DtoEnum受払区分) TableName() string {
	return "Enum受払区分"
}
func (d DtoEnum受払区分) RowState() DataRowState {
	return d.rowState
}

// Import はDtoEnum受払区分型に主キー以外を上書きする。
func (d *DtoEnum受払区分) Import(名称 string, 受入フラグ bool, 符号 int) {
	// 項目がすべて一致していたら、何もしない
	if d.Fld名称 == 名称 && d.Fld受入フラグ == 受入フラグ && d.Fld符号 == 符号 {
		return
	}
	if d.Fld名称 != 名称 {
		d.Fld名称 = 名称
		d.Ub.Set(TblEnum受払区分().Fld名称(), 名称)
	}
	if d.Fld受入フラグ != 受入フラグ {
		d.Fld受入フラグ = 受入フラグ
		d.Ub.Set(TblEnum受払区分().Fld受入フラグ(), 受入フラグ)
	}
	if d.Fld符号 != 符号 {
		d.Fld符号 = 符号
		d.Ub.Set(TblEnum受払区分().Fld符号(), 符号)
	}

}
