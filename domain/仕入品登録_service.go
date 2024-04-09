package domain

import (
	"errors"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"

	"golang.org/x/xerrors"
)

type Srv仕入品登録 struct {
	rm objects.RepManager
}

func NewSrv仕入品登録(rm objects.RepManager) *Srv仕入品登録 {
	return &Srv仕入品登録{
		rm: rm,
	}
}

// Exec登録 は仕入品を新規登録または上書きします。
func (s *Srv仕入品登録) Exec登録(
	アップロード履歴 types.No,
	コード types.Code品目,
	名称 string,
	基準単位コード types.Code単位,
	生産用品目区分コード types.Code生産用品目区分,
	標準単価 types.Price,
) error {
	rep品目 := s.rm.NewRep品目()
	rep単位 := s.rm.NewRep単位()
	rep生産用品目区分 := s.rm.NewRep生産用品目区分()

	基準単位, notFound := rep単位.GetBy(基準単位コード)
	if notFound != nil {
		return notFound
	}

	生産用品目区分, notFound := rep生産用品目区分.GetBy(生産用品目区分コード)
	if notFound != nil {
		return notFound
	}

	if 既存e, notFound := rep品目.Get仕入品By(コード); notFound == nil {
		既存e.Getコード = コード
		既存e.Get名称 = 名称
		既存e.Get基準単位 = 基準単位
		既存e.Get生産用品目区分 = 生産用品目区分
		既存e.Get標準単価 = 標準単価
	} else if errors.Is(notFound, types.ErrNotFound) {
		var e *objects.Ent品目仕入品
		if 既存品目e, notFound := rep品目.Get品目By(コード); notFound == nil {
			既存品目e.Get基準単位 = 基準単位
			既存品目e.Get生産用品目区分 = 生産用品目区分
			var err error
			e, err = objects.NewEnt品目仕入品(既存品目e, 標準単価)
			if err != nil {
				return err
			}
		} else if errors.Is(notFound, types.ErrNotFound) {
			var err error
			品目e, err := objects.NewEnt品目(
				コード,
				名称,
				基準単位,
				生産用品目区分,
			)
			if err != nil {
				return err
			}

			e, err = objects.NewEnt品目仕入品(品目e, 標準単価)
			if err != nil {
				return err
			}
		} else {
			return xerrors.Errorf(" :%w", notFound)
		}

		rep品目.AddNew仕入品(e)
	} else {
		return xerrors.Errorf(" :%w", notFound)
	}
	rep品目.Save(アップロード履歴)
	return nil
}
