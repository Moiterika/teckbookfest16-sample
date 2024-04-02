// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"encoding/json"
	objects "techbookfest16-sample/domain/objects"
)

type Dto生産用品目区分 struct {
	FldID      Id                  `json:"ID"`
	Fldコード     objects.Code生産用品目区分 `json:"コード"`
	Fld名称      string              `json:"名称"`
	Fld何かのフラグ1 bool                `json:"何かのフラグ1"`
	Fld何かのフラグ2 bool                `json:"何かのフラグ2"`

	rowState DataRowState
	Ub       *ub生産用品目区分 `json:"-"`
}

func (d Dto生産用品目区分) TableName() string {
	return "生産用品目区分"
}
func (d Dto生産用品目区分) RowState() DataRowState {
	return d.rowState
}

// Import はDto生産用品目区分型に主キー以外を上書きする。
func (d *Dto生産用品目区分) Import(コード objects.Code生産用品目区分, 名称 string, 何かのフラグ1 bool, 何かのフラグ2 bool) {
	// 項目がすべて一致していたら、何もしない
	if d.Fldコード == コード && d.Fld名称 == 名称 && d.Fld何かのフラグ1 == 何かのフラグ1 && d.Fld何かのフラグ2 == 何かのフラグ2 {
		return
	}
	if d.Fldコード != コード {
		d.Fldコード = コード
		d.Ub.Set(Tbl生産用品目区分().Fldコード(), コード)
	}
	if d.Fld名称 != 名称 {
		d.Fld名称 = 名称
		d.Ub.Set(Tbl生産用品目区分().Fld名称(), 名称)
	}
	if d.Fld何かのフラグ1 != 何かのフラグ1 {
		d.Fld何かのフラグ1 = 何かのフラグ1
		d.Ub.Set(Tbl生産用品目区分().Fld何かのフラグ1(), 何かのフラグ1)
	}
	if d.Fld何かのフラグ2 != 何かのフラグ2 {
		d.Fld何かのフラグ2 = 何かのフラグ2
		d.Ub.Set(Tbl生産用品目区分().Fld何かのフラグ2(), 何かのフラグ2)
	}

}

// jsonKey生産用品目区分 はロギング用jsonのキー。主キー項目
type jsonKey生産用品目区分 struct {
	FldID Id `json:"ID"`
}

// jsonKey はロギング用jsonのキーを生成するメソッド。
func (d *Dto生産用品目区分) jsonKey() jsonKey生産用品目区分 {
	return jsonKey生産用品目区分{FldID: d.FldID}
}

type json生産用品目区分 struct {
	RowState string         `json:"row_state"`
	K        jsonKey生産用品目区分 `json:"k"`
	V        *Dto生産用品目区分    `json:"v,omitempty"`
}

// ToJson はMarshalJSONと同じ機能を提供するメソッド。しかし、無限ループを防ぐため、別名メソッドにしてある。
func (d *Dto生産用品目区分) ToJson() ([]byte, error) {
	switch d.rowState {
	case Modified:
		j := make(map[string]interface{})
		j["row_state"] = d.rowState.String()
		j["k"] = map[string]interface{}{"ID": d.FldID}
		j["v"] = d.Ub.copyMap()
		return json.Marshal(j)
	case Deleted:
		return json.Marshal(json生産用品目区分{
			K:        d.jsonKey(),
			RowState: d.rowState.String(),
			V:        nil,
		})
	default:
		return json.Marshal(json生産用品目区分{
			K:        d.jsonKey(),
			RowState: d.rowState.String(),
			V:        d,
		})
	}
}
