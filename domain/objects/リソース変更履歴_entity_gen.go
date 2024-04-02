// Code generated by xlsx2go.go; DO NOT EDIT.

package objects

import "time"

type Entリソース変更履歴 struct {
	GetID      No                  `json:"ID"`
	Get登録日時    time.Time           `json:"登録日時"`
	Getリソース名   string              `json:"リソース名"`
	Get変更区分    string              `json:"変更区分"`
	Get変更内容    []byte              `json:"変更内容"`
	Getアップロード時 *Valリソース変更履歴アップロード時 `json:"アップロード時,omitempty"`
}

func NewEntリソース変更履歴(ID No, 登録日時 time.Time, リソース名 string, 変更区分 string, 変更内容 []byte, アップロード時 *Valリソース変更履歴アップロード時) (*Entリソース変更履歴, error) {
	e := &Entリソース変更履歴{
		GetID:      ID,
		Getアップロード時: アップロード時,
		Getリソース名:   リソース名,
		Get変更内容:    変更内容,
		Get変更区分:    変更区分,
		Get登録日時:    登録日時,
	}
	err := e.Validate()
	return e, err
}
