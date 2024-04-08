package infra

import (
	"errors"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra/dao"

	"golang.org/x/xerrors"
)

type repTrn単位 struct {
	rm       *repManagerTrn
	isLoaded bool
	//objects.Rep単位
}

func (r *repTrn単位) init() error {
	dt単位, err := r.rm.dm.NewDaoTrn単位().Dt()
	if err != nil {
		return xerrors.Errorf(" :%w", err)
	}
	r.rm.list単位 = make([]*objects.Ent単位, len(dt単位))
	for i, dr := range dt単位 {
		e, err := objects.NewEnt単位(
			dr.Fldコード,
			dr.Fld名称,
		)
		if err != nil {
			return xerrors.Errorf(" :%w", err)
		}

		r.rm.list単位[i] = e
		r.rm.mapIDvs単位[dr.FldID] = e
		r.rm.mapコードvs単位[e.Getコード] = e
	}
	r.isLoaded = true
	return nil
}

func (r *repTrn単位) List() ([]*objects.Ent単位, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			xerrors.Errorf(" :%w", err)
			return nil, err
		}
	}
	return r.rm.list単位, nil
}

func (r *repTrn単位) getBy(id dao.Id) (*objects.Ent単位, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapIDvs単位[id]
	if !ok {
		return nil, xerrors.Errorf("単位が見つかりません。単位ID=%s: %w", id, objects.ErrNotFound)
	}
	return e, nil
}

func (r *repTrn単位) GetBy(コード types.Code単位) (*objects.Ent単位, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapコードvs単位[コード]
	if !ok {
		return nil, xerrors.Errorf("単位が見つかりません。単位コード=%s: %w", コード, objects.ErrNotFound)
	}
	return e, nil
}

func (r *repTrn単位) AddNew(e *objects.Ent単位) error {
	// エンティティの責務ではなく、コレクション重複チェックはリポジトリーの責務とする
	if _, notFound := r.GetBy(e.Getコード); !errors.Is(notFound, objects.ErrNotFound) {
		return xerrors.Errorf("単位がすでに存在します。単位コード=%s: %w", e.Getコード, objects.ErrAlreadyExists)
	}

	r.rm.list単位 = append(r.rm.list単位, e)

	return nil
}

// TODO リソースの一括アップロードやイベント系で1行ずつロギングしない場合はMultiInsertを使う
func (r *repTrn単位) Save(アップロード履歴ID types.No) error {
	dao単位 := r.rm.dm.NewDaoTrn単位()
	defer dao単位.Reset()
	logger := newCmdTrnリソース変更履歴(r.rm.dm)

	for _, e := range r.rm.list単位 {
		if dr, notFound := dao単位.GetByCode(e.Getコード); !errors.Is(notFound, dao.NotFoundError) {
			dr.Import(e.Getコード, e.Get名称)
			_, err := dao単位.UpdateBy(dr)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			logger.Write(dr, アップロード履歴ID)
		} else {
			dr := &dao.Dto単位{
				Fldコード: e.Getコード,
				Fld名称:  e.Get名称,
				Ub:     dao.NewUb単位(),
			}
			err := dao単位.Insert(dr)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			logger.Write(dr, アップロード履歴ID)
		}
	}

	return nil
}
