package domain

import (
	"errors"
	"fmt"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"

	"golang.org/x/xerrors"
)

type Srv単位登録 struct {
	rep objects.Rep単位
}

func NewSrv単位登録(rep objects.Rep単位) *Srv単位登録 {
	return &Srv単位登録{
		rep: rep,
	}
}

// Exec登録 は単位を新規登録または上書きします。
func (s *Srv単位登録) Exec登録(アップロード履歴 types.No, e *objects.Ent単位) error {
	err := e.Validate()
	if err != nil {
		return fmt.Errorf("validate error: %w, %w", err, types.ErrArg)
	}

	if 既存e, notFound := s.rep.GetBy(e.Getコード); notFound == nil {
		既存e.Getコード = e.Getコード
		既存e.Get名称 = e.Get名称
	} else if errors.Is(notFound, types.ErrNotFound) {
		s.rep.AddNew(e)
	} else {
		return xerrors.Errorf(" :%w", notFound)
	}
	s.rep.Save(アップロード履歴)
	return nil
}
