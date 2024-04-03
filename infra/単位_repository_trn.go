package infra

import (
	"errors"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra/dao"

	"github.com/Moiterika/a"

	"golang.org/x/xerrors"
)

type repTrn単位 struct {
	rm *repManagerTrn
	objects.Rep単位
}

func (r *repTrn単位) init() (err error) {
	r.rm.list単位, err = r.list()
	if err != nil {
		xerrors.Errorf(" :%w", err)
		return
	}
	r.rm.mapIDvs単位 = a.ToMap(r.rm.list単位, func(e *objects.Ent単位) dao.Id {
		dr, _ := r.rm.dm.NewDaoTrn単位().GetByCode(e.Getコード)
		return dr.FldID
	})
	r.rm.mapコードvs単位 = a.ToMap(r.rm.list単位, func(e *objects.Ent単位) types.Code単位 {
		return e.Getコード
	})
	return
}

func (r *repTrn単位) list() (list []*objects.Ent単位, err error) {
	dt単位, errDt := r.rm.dm.NewDaoTrn単位().Dt()
	if errDt != nil {
		err = xerrors.Errorf(" :%w", errDt)
		return
	}
	list = make([]*objects.Ent単位, len(dt単位))
	for i, dr := range dt単位 {
		e := &objects.Ent単位{
			Getコード: dr.Fldコード,
			Get名称:  dr.Fld名称,
		}
		list[i] = e
	}
	return
}

func (r *repTrn単位) List() ([]*objects.Ent単位, error) {
	if len(r.rm.list単位) == 0 {
		err := r.init()
		if err != nil {
			xerrors.Errorf(" :%w", err)
			return nil, err
		}
	}
	return r.rm.list単位, nil
}

func (r *repTrn単位) getBy(id dao.Id) (e *objects.Ent単位, err error) {
	if len(r.rm.mapIDvs単位) == 0 {
		err = r.init()
		if err != nil {
			xerrors.Errorf(" :%w", err)
			return
		}
	}
	var ok bool
	e, ok = r.rm.mapIDvs単位[id]
	if !ok {
		err = xerrors.Errorf("単位が見つかりません。単位ID=%s: %w", id, objects.ErrNotFound)
		return
	}
	return
}

func (r *repTrn単位) GetBy(コード types.Code単位) (e *objects.Ent単位, err error) {
	if len(r.rm.mapIDvs単位) == 0 {
		err = r.init()
		if err != nil {
			xerrors.Errorf(" :%w", err)
			return
		}
	}
	var ok bool
	e, ok = r.rm.mapコードvs単位[コード]
	if !ok {
		err = xerrors.Errorf("単位が見つかりません。単位コード=%s: %w", コード, objects.ErrNotFound)
		return
	}
	return
}

func (r *repTrn単位) AddNew(e *objects.Ent単位) error {
	// エンティティの責務ではなく、コレクション重複チェックはリポジトリーの責務とする
	if _, err := r.GetBy(e.Getコード); !errors.Is(err, objects.ErrNotFound) {
		return xerrors.Errorf("単位コードがすでに存在します。単位コード=%s: %w", e.Getコード, objects.ErrAlreadyExists)
	}

	r.rm.list単位 = append(r.rm.list単位, e)

	return nil
}

// TODO リソースの一括アップロードやイベント系で1行ずつロギングしない場合はMultiInsertを使う
func (r *repTrn単位) Save(アップロード履歴ID types.No) (err error) {
	dao単位 := r.rm.dm.NewDaoTrn単位()
	defer dao単位.Reset()
	logger := newCmdTrnリソース変更履歴(r.rm.dm)

	for _, e := range r.rm.list単位 {
		if dr, errCode := dao単位.GetByCode(e.Getコード); !errors.Is(errCode, dao.NotFoundError) {
			dr.Import(e.Getコード, e.Get名称)
			_, err = dao単位.UpdateBy(dr)
			if err != nil {
				err = xerrors.Errorf(": %w", err)
				return
			}
			logger.Write(dr, アップロード履歴ID)
		} else {
			dr := &dao.Dto単位{
				Fldコード: e.Getコード,
				Fld名称:  e.Get名称,
				Ub:     dao.NewUb単位(),
			}
			dr.FldID, err = dao単位.Insert(dr)
			if err != nil {
				err = xerrors.Errorf(": %w", err)
				return
			}
			logger.Write(dr, アップロード履歴ID)
		}
	}

	return
}
