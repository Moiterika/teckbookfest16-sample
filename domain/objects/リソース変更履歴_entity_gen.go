// Code generated by xlsx2go.go; DO NOT EDIT.

package objects

import (
	types "techbookfest16-sample/domain/types"
	"time"
)

type Entリソース変更履歴 struct {
	GetNo    types.No  `json:"No"`
	Get登録日時  time.Time `json:"登録日時"`
	Getリソース名 string    `json:"リソース名"`
	Get変更区分  string    `json:"変更区分"`
	Get変更内容  []byte    `json:"変更内容"`
}

func NewEntリソース変更履歴(No types.No, 登録日時 time.Time, リソース名 string, 変更区分 string, 変更内容 []byte) (*Entリソース変更履歴, error) {
	e := &Entリソース変更履歴{
		GetNo:    No,
		Getリソース名: リソース名,
		Get変更内容:  変更内容,
		Get変更区分:  変更区分,
		Get登録日時:  登録日時,
	}
	err := e.Validate()
	return e, err
}
func (e *Entリソース変更履歴) Id() types.No {
	return e.GetNo
}
func (e *Entリソース変更履歴) Equals(other types.Identifier[types.No]) bool {
	if other == nil {
		return false
	}
	if other.Equals(e) {
		return true
	}
	return false
}
