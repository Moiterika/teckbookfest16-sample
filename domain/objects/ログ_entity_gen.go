// Code generated by xlsx2go.go; DO NOT EDIT.

package objects

import (
	types "techbookfest16-sample/domain/types"
	"time"
)

type Entログ struct {
	GetNo   types.No  `json:"No"`
	Get登録日時 time.Time `json:"登録日時"`
	Get区分   Enumログ区分  `json:"区分"`
	Get内容   string    `json:"内容"`
}

func NewEntログ(No types.No, 登録日時 time.Time, 区分 Enumログ区分, 内容 string) (*Entログ, error) {
	e := &Entログ{
		GetNo:   No,
		Get内容:   内容,
		Get区分:   区分,
		Get登録日時: 登録日時,
	}
	err := e.Validate()
	return e, err
}
func (e *Entログ) Id() types.No {
	return e.GetNo
}
func (e *Entログ) Equals(other types.Identifier[types.No]) bool {
	if other == nil {
		return false
	}
	if other.Equals(e) {
		return true
	}
	return false
}
