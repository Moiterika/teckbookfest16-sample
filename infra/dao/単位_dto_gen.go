// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"encoding/json"
	types "techbookfest16-sample/domain/types"
)

type Dto単位 struct {
	FldID  Id           `json:"ID"`
	Fldコード types.Code単位 `json:"コード"`
	Fld名称  string       `json:"名称"`

	rowState DataRowState
	Ub       *ub単位 `json:"-"`
}

func (d Dto単位) TableName() string {
	return "単位"
}
func (d Dto単位) RowState() DataRowState {
	return d.rowState
}

// Import はDto単位型に主キー以外を上書きする。
func (d *Dto単位) Import(コード types.Code単位, 名称 string) {
	// 項目がすべて一致していたら、何もしない
	if d.Fldコード == コード && d.Fld名称 == 名称 {
		return
	}
	if d.Fldコード != コード {
		d.Fldコード = コード
		d.Ub.Set(Tbl単位().Fldコード(), コード)
	}
	if d.Fld名称 != 名称 {
		d.Fld名称 = 名称
		d.Ub.Set(Tbl単位().Fld名称(), 名称)
	}

}

// jsonKey単位 はロギング用jsonのキー。主キー項目
type jsonKey単位 struct {
	FldID Id `json:"ID"`
}

// jsonKey はロギング用jsonのキーを生成するメソッド。
func (d *Dto単位) jsonKey() jsonKey単位 {
	return jsonKey単位{FldID: d.FldID}
}

type json単位 struct {
	RowState string    `json:"row_state"`
	K        jsonKey単位 `json:"k"`
	V        *Dto単位    `json:"v,omitempty"`
}

// ToJson はMarshalJSONと同じ機能を提供するメソッド。しかし、無限ループを防ぐため、別名メソッドにしてある。
func (d *Dto単位) ToJson() ([]byte, error) {
	switch d.rowState {
	case Modified:
		j := make(map[string]interface{})
		j["row_state"] = d.rowState.String()
		j["k"] = map[string]interface{}{"ID": d.FldID}
		j["v"] = d.Ub.copyMap()
		return json.Marshal(j)
	case Deleted:
		return json.Marshal(json単位{
			K:        d.jsonKey(),
			RowState: d.rowState.String(),
			V:        nil,
		})
	default:
		return json.Marshal(json単位{
			K:        d.jsonKey(),
			RowState: d.rowState.String(),
			V:        d,
		})
	}
}
