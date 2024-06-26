// Code generated by xlsx2go.go; DO NOT EDIT.

package objects

import (
	types "techbookfest16-sample/domain/types"
	"time"
)

type Ent受払 struct {
	GetNo    types.No        `json:"No"`
	Get登録日時  time.Time       `json:"登録日時"`
	Get計上月   time.Time       `json:"計上月"`
	Get受払区分  Enum受払区分        `json:"受払区分"`
	Get赤伝フラグ bool            `json:"赤伝フラグ"`
	Get品目    *Ent品目          `json:"品目,omitempty"`
	Get基準数量  types.Inventory `json:"基準数量"`
}

func NewEnt受払(登録日時 time.Time, 計上月 time.Time, 受払区分 Enum受払区分, 赤伝フラグ bool, 品目 *Ent品目, 基準数量 types.Inventory) (*Ent受払, error) {
	e := &Ent受払{
		Get受払区分:  受払区分,
		Get品目:    品目,
		Get基準数量:  基準数量,
		Get登録日時:  登録日時,
		Get計上月:   計上月,
		Get赤伝フラグ: 赤伝フラグ,
	}
	err := e.Validate()
	return e, err
}
