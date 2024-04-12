package infra

import (
	"errors"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra/dao"

	"golang.org/x/xerrors"
)

type repTrn品目 struct {
	rm       *repManagerTrn
	isLoaded bool
	// 品目tempId  dao.Id
	// 仕入品tempId dao.Id
	wb品目  dao.Wb品目
	wb仕入品 dao.Wb品目仕入品
	wb製造品 dao.Wb品目製造品
}

func (r *repTrn品目) init() error {
	dt品目, err := r.rm.dm.NewDaoTrn品目With(r.wb品目).Dt()
	if err != nil {
		return xerrors.Errorf(" :%w", err)
	}

	rep単位 := r.rm.rep単位()
	rep生産用品目区分 := r.rm.rep生産用品目区分()

	r.rm.list品目 = make([]*objects.Ent品目, len(dt品目))
	for i, dr := range dt品目 {
		基準単位, err := rep単位.getBy(dr.Fld基準単位ID)
		if err != nil {
			return xerrors.Errorf("%s :%w", err.Error(), ErrInit)
		}

		生産用品目区分, err := rep生産用品目区分.getBy(dr.Fld生産用品目区分ID)
		if err != nil {
			return xerrors.Errorf("%s :%w", err.Error(), ErrInit)
		}

		e := &objects.Ent品目{
			Getコード:     types.Code品目(dr.Fldコード),
			Get名称:      dr.Fld名称,
			Get基準単位:    基準単位,
			Get生産用品目区分: 生産用品目区分,
		}
		r.rm.list品目[i] = e
		r.rm.mapIDvs品目[dr.FldID] = e
		r.rm.mapコードvs品目[e.Getコード] = e
	}

	dt仕入品, err := r.rm.dm.NewDaoTrn品目仕入品With(r.wb仕入品).Dt()
	if err != nil {
		return xerrors.Errorf("%s :%w", err.Error(), ErrInit)
	}
	r.rm.list仕入品 = make([]*objects.Ent品目仕入品, len(dt仕入品))
	for i, dr := range dt仕入品 {
		// r.rm.mapIDvs品目を直接参照する。
		// 理由：この時点では!r.isLoadedなので、rep品目.getByを呼び出すと無限ループするため。
		e品目, ok := r.rm.mapIDvs品目[dr.FldID]
		if !ok {
			return xerrors.Errorf("品目が見つかりません。品目ID=%d :%w", dr.FldID, ErrInit)
		}
		e単位, err := rep単位.getBy(dr.Fld標準単価単位ID)
		if err != nil {
			return xerrors.Errorf("%s :%w", err.Error(), ErrInit)
		}
		e := &objects.Ent品目仕入品{
			Ent品目:   e品目,
			Get標準単価: types.NewPrice(dr.Fld標準単価, dr.Fld標準単価通貨ID, e単位.Getコード),
		}
		r.rm.list仕入品[i] = e
		r.rm.mapIDvs仕入品[dr.FldID] = e
		r.rm.mapコードvs仕入品[e.Getコード] = e
	}
	r.isLoaded = true
	return nil
}

func (r *repTrn品目) Get品目一覧() ([]*objects.Ent品目, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}

	return r.rm.list品目, nil
}

func (r *repTrn品目) Get仕入品一覧() ([]*objects.Ent品目仕入品, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}

	return r.rm.list仕入品, nil
}

func (r *repTrn品目) getBy(id dao.Id) (*objects.Ent品目, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapIDvs品目[id]
	if !ok {
		return nil, xerrors.Errorf("品目が見つかりません。品目ID=%d: %w", id, types.ErrNotFound)
	}
	return e, nil
}

func (r *repTrn品目) Get品目By(コード types.Code品目) (*objects.Ent品目, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapコードvs品目[コード]
	if !ok {
		return nil, xerrors.Errorf("品目が見つかりません。品目コード=%s: %w", コード, types.ErrNotFound)
	}
	return e, nil
}

func (r *repTrn品目) Get仕入品By(コード types.Code品目) (*objects.Ent品目仕入品, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapコードvs仕入品[コード]
	if !ok {
		return nil, xerrors.Errorf("品目が見つかりません。品目コード=%s: %w", コード, types.ErrNotFound)
	}
	return e, nil
}

func (r *repTrn品目) AddNew仕入品(e *objects.Ent品目仕入品) error {
	// エンティティの責務ではなく、コレクション重複チェックはリポジトリーの責務とする
	if _, err := r.Get仕入品By(e.Getコード); !errors.Is(err, types.ErrNotFound) {
		return xerrors.Errorf("仕入品がすでに存在します。品目コード=%s: %w", e.Getコード, types.ErrAlreadyExists)
	}

	if _, err := r.Get品目By(e.Getコード); errors.Is(err, types.ErrNotFound) {
		r.rm.list品目 = append(r.rm.list品目, e.Ent品目)
		// r.品目tempId--
		// r.rm.mapIDvs品目[r.品目tempId] = e.Ent品目
		r.rm.mapコードvs品目[e.Getコード] = e.Ent品目
	}

	r.rm.list仕入品 = append(r.rm.list仕入品, e)
	// r.仕入品tempId--
	// r.rm.mapIDvs仕入品[r.仕入品tempId] = e
	r.rm.mapコードvs仕入品[e.Getコード] = e
	return nil
}

// TODO リソースの一括アップロードやイベント系で1行ずつロギングしない場合はMultiInsertを使う
func (r *repTrn品目) Save(アップロード履歴ID types.No) error {
	dao品目 := r.rm.dm.NewDaoTrn品目()
	dao仕入品 := r.rm.dm.NewDaoTrn品目仕入品()
	defer dao品目.Reset()
	defer dao仕入品.Reset()
	logger := newCmdTrnリソース変更履歴(r.rm.dm)

	dao単位 := r.rm.dm.NewDaoTrn単位()
	dao生産用品目区分 := r.rm.dm.NewDaoTrn生産用品目区分()

	for _, e := range r.rm.list品目 {
		dr単位, err := dao単位.GetByCode(e.Get基準単位.Getコード)
		if err != nil {
			return xerrors.Errorf("品目コード=%s: %w", e.Getコード, err)
		}
		dr生産用品目区分, err := dao生産用品目区分.GetByCode(e.Get生産用品目区分.Getコード)
		if err != nil {
			return xerrors.Errorf("品目コード=%s: %w", e.Getコード, err)
		}

		if dr, notFound := dao品目.GetByCode(e.Getコード); errors.Is(notFound, types.ErrNotFound) {
			dr := &dao.Dto品目{
				//FldID:        0,
				Fldコード:       e.Getコード,
				Fld名称:        e.Get名称,
				Fld基準単位ID:    dr単位.FldID,
				Fld生産用品目区分ID: dr生産用品目区分.FldID,
				Ub:           dao.NewUb品目(),
			}
			err = dao品目.Insert(dr)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			logger.Write(dr, アップロード履歴ID)
		} else {
			dr.Import(e.Getコード, e.Get名称, dr単位.FldID, dr生産用品目区分.FldID)
			_, err = dao品目.UpdateBy(dr)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			logger.Write(dr, アップロード履歴ID)
		}
	}

	for _, e := range r.rm.list仕入品 {
		dr品目, err := dao品目.GetByCode(e.Getコード)
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}
		dr単位, err := dao単位.GetByCode(e.Get基準単位.Getコード)
		if err != nil {
			return xerrors.Errorf("品目コード=%s: %w", e.Getコード, err)
		}
		if dr, notFound := dao仕入品.GetBy(dr品目.FldID); errors.Is(notFound, types.ErrNotFound) {
			dr := &dao.Dto品目仕入品{
				FldID:       dr品目.FldID,
				Fld標準単価:     e.Get標準単価.Amt(),
				Fld標準単価通貨ID: e.Get標準単価.Cur(),
				Fld標準単価単位ID: dr単位.FldID,
				Ub:          dao.NewUb品目仕入品(),
			}
			err = dao仕入品.Insert(dr)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			logger.Write(dr, アップロード履歴ID)
		} else {
			dr.Import(e.Get標準単価.Amt(),
				e.Get標準単価.Cur(),
				dr単位.FldID)
			_, err = dao仕入品.UpdateBy(dr)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			logger.Write(dr, アップロード履歴ID)
		}
	}

	return nil
}
