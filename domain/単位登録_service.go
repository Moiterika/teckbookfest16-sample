package domain

import (
	"errors"
	"fmt"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"

	"golang.org/x/xerrors"
)

type Srv単位登録 struct {
	rm objects.RepManager
}

func NewSrv単位登録(rm objects.RepManager) *Srv単位登録 {
	return &Srv単位登録{
		rm: rm,
	}
}

// Exec登録 は単位を新規登録または上書きします。
func (s *Srv単位登録) Exec登録(アップロード履歴 types.No, e *objects.Ent単位) (err error) {
	err = e.Validate()
	if err != nil {
		err = fmt.Errorf("validate error: %w, %w", err, objects.ErrArg)
		return
	}

	rep := s.rm.NewRep単位()
	if 既存単位, err既存単位 := rep.GetBy(e.Getコード); err既存単位 == nil {
		既存単位.Getコード = e.Getコード
		既存単位.Get名称 = e.Get名称
	} else if errors.Is(err既存単位, objects.ErrNotFound) {
		rep.AddNew(e)
	} else {
		err = xerrors.Errorf(" :%w", err既存単位)
		return
	}
	rep.Save(アップロード履歴)
	return
}
