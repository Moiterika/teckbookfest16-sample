// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import objects "techbookfest16-sample/domain/objects"

type DtoEnumログ区分 struct {
	Fld区分 objects.Enumログ区分
	Fld名称 string

	rowState DataRowState
	Ub       *ubEnumログ区分 `json:"-"`
}

func (d DtoEnumログ区分) TableName() string {
	return "Enumログ区分"
}
func (d DtoEnumログ区分) RowState() DataRowState {
	return d.rowState
}

// Import はDtoEnumログ区分型に主キー以外を上書きする。
func (d *DtoEnumログ区分) Import(名称 string) {
	// 項目がすべて一致していたら、何もしない
	if d.Fld名称 == 名称 {
		return
	}
	if d.Fld名称 != 名称 {
		d.Fld名称 = 名称
		d.Ub.Set(TblEnumログ区分().Fld名称(), 名称)
	}

}
