package infra

import (
	"errors"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra/dao"

	"golang.org/x/xerrors"
)

type repTrn生産用品目区分 struct {
	rm       *repManagerTrn
	isLoaded bool
}

func (r *repTrn生産用品目区分) init() error {
	dt生産用品目区分, err := r.rm.dm.NewDaoTrn生産用品目区分().Dt()
	if err != nil {
		return xerrors.Errorf(" :%w", err)
	}
	r.rm.list生産用品目区分 = make([]*objects.Ent生産用品目区分, len(dt生産用品目区分))
	for i, dr := range dt生産用品目区分 {
		e, err := objects.NewEnt生産用品目区分(
			types.Code生産用品目区分(dr.Fldコード),
			dr.Fld名称,
			dr.Fld何かのフラグ1,
			dr.Fld何かのフラグ2,
		)
		if err != nil {
			return xerrors.Errorf(" :%w", err)
		}
		r.rm.list生産用品目区分[i] = e
		r.rm.mapIDvs生産用品目区分[dr.FldID] = e
		r.rm.mapコードvs生産用品目区分[e.Getコード] = e
	}

	r.isLoaded = true

	return nil
}

func (r *repTrn生産用品目区分) List() ([]*objects.Ent生産用品目区分, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	return r.rm.list生産用品目区分, nil
}

func (r *repTrn生産用品目区分) getBy(id dao.Id) (*objects.Ent生産用品目区分, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapIDvs生産用品目区分[id]
	if !ok {
		return nil, xerrors.Errorf("生産用品目区分が見つかりません。生産用品目区分ID=%s: %w", id, objects.ErrNotFound)
	}
	return e, nil
}

func (r *repTrn生産用品目区分) GetBy(コード types.Code生産用品目区分) (*objects.Ent生産用品目区分, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapコードvs生産用品目区分[コード]
	if !ok {
		return nil, xerrors.Errorf("生産用品目区分が見つかりません。生産用品目区分コード=%s: %w", コード, objects.ErrNotFound)
	}
	return e, nil
}

func (r *repTrn生産用品目区分) AddNew(e *objects.Ent生産用品目区分) error {
	// エンティティの責務ではなく、コレクション重複チェックはリポジトリーの責務とする
	if _, notFound := r.GetBy(e.Getコード); !errors.Is(notFound, objects.ErrNotFound) {
		return xerrors.Errorf("生産用品目区分がすでに存在します。生産用品目区分コード=%s: %w", e.Getコード, objects.ErrAlreadyExists)
	}

	r.rm.list生産用品目区分 = append(r.rm.list生産用品目区分, e)

	return nil
}

// TODO リソースの一括アップロードやイベント系で1行ずつロギングしない場合はMultiInsertを使う
func (r *repTrn生産用品目区分) Save(アップロード履歴ID types.No) (err error) {
	dao生産用品目区分 := r.rm.dm.NewDaoTrn生産用品目区分()
	defer dao生産用品目区分.Reset()
	logger := newCmdTrnリソース変更履歴(r.rm.dm)

	for _, e := range r.rm.list生産用品目区分 {
		if dr, notFound := dao生産用品目区分.GetByCode(e.Getコード); !errors.Is(notFound, dao.NotFoundError) {
			dr.Import(e.Getコード, e.Get名称, e.Get何かのフラグ1, e.Get何かのフラグ2)
			_, err = dao生産用品目区分.UpdateBy(dr)
			if err != nil {
				err = xerrors.Errorf(": %w", err)
				return
			}
			logger.Write(dr, アップロード履歴ID)
		} else {
			dr := &dao.Dto生産用品目区分{
				Fldコード:     e.Getコード,
				Fld名称:      e.Get名称,
				Fld何かのフラグ1: e.Get何かのフラグ1,
				Fld何かのフラグ2: e.Get何かのフラグ2,
				Ub:         dao.NewUb生産用品目区分(),
			}
			err = dao生産用品目区分.Insert(dr)
			if err != nil {
				err = xerrors.Errorf(": %w", err)
				return
			}
			logger.Write(dr, アップロード履歴ID)
		}
	}

	return
}
