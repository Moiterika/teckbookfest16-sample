package infra

import (
	"errors"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra/dao"

	"github.com/Moiterika/a"

	"golang.org/x/xerrors"
)

type repTrn品目 struct {
	rm *repManagerTrn
}

func (r *repTrn品目) init() (err error) {
	r.rm.list品目, err = r.list()
	if err != nil {
		xerrors.Errorf(" :%w", err)
		return
	}
	r.rm.mapIDvs品目 = a.ToMap(r.rm.list品目, func(e *objects.Ent品目) dao.Id {
		dr, _ := r.rm.dm.NewDaoTrn品目().GetByCode(e.Getコード)
		return dr.FldID
	})
	r.rm.mapコードvs品目 = a.ToMap(r.rm.list品目, func(e *objects.Ent品目) objects.Code品目 {
		return e.Getコード
	})
	return
}

func (r *repTrn品目) list() (list []*objects.Ent品目, err error) {
	dt品目, errDt := r.rm.dm.NewDaoTrn品目().Dt()
	if errDt != nil {
		err = xerrors.Errorf(" :%w", errDt)
		return
	}

	rep単位 := r.rm.rep単位()
	rep生産用品目区分 := r.rm.rep生産用品目区分()

	dao仕入品 := r.rm.dm.NewDaoTrn品目仕入品()
	dao製造品 := r.rm.dm.NewDaoTrn品目製造品()

	list = make([]*objects.Ent品目, len(dt品目))
	for i, dr := range dt品目 {
		基準単位, err基準単位 := rep単位.Get(dr.Fld基準単位ID)
		if err基準単位 != nil {
			err = xerrors.Errorf(" :%w", err基準単位)
			return
		}
		生産用品目区分, err生産用品目区分 := rep生産用品目区分.getBy(dr.Fld生産用品目区分ID)
		if err生産用品目区分 != nil {
			err = xerrors.Errorf(" :%w", err生産用品目区分)
			return
		}
		dr仕入品, err仕入品 := dao仕入品.GetBy(dr.FldID)
		if err仕入品 != nil {
			err = xerrors.Errorf(" :%w", err仕入品)
			return
		}
		var 仕入品 *objects.Val品目仕入品
		if dr仕入品 != nil {
			仕入品標準単価単位, err仕入品標準単価単位 := rep単位.Get(dr仕入品.Fld標準単価単位ID)
			if err仕入品標準単価単位 != nil {
				err = xerrors.Errorf(" :%w", err仕入品標準単価単位)
				return
			}
			仕入品 = &objects.Val品目仕入品{
				Get標準単価: types.NewPrice(dr仕入品.Fld標準単価, dr仕入品.Fld標準単価通貨ID, 仕入品標準単価単位.GetID),
			}
		}
		dr製造品, err製造品 := dao製造品.GetBy(dr.FldID)
		if err製造品 != nil {
			err = xerrors.Errorf(" :%w", err製造品)
			return
		}
		var 製造品 *objects.Val品目製造品
		if dr製造品 != nil {
			製造品 = &objects.Val品目製造品{
				GetMRP計算対象フラグ: dr製造品.FldMRP計算対象フラグ,
			}
		}

		e := &objects.Ent品目{
			Getコード:     objects.Code品目(dr.Fldコード),
			Get名称:      dr.Fld名称,
			Get基準単位:    基準単位,
			Get生産用品目区分: 生産用品目区分,
			Get仕入品:     仕入品,
			Get製造品:     製造品,
		}
		list[i] = e
	}
	return
}

func (r *repTrn品目) List() ([]*objects.Ent品目, error) {
	if len(r.rm.list品目) == 0 {
		err := r.init()
		if err != nil {
			xerrors.Errorf(" :%w", err)
			return nil, err
		}
	}
	return r.rm.list品目, nil
}

func (r *repTrn品目) getBy(id dao.Id) (e *objects.Ent品目, err error) {
	if len(r.rm.mapIDvs品目) == 0 {
		err = r.init()
		if err != nil {
			xerrors.Errorf(" :%w", err)
			return
		}
	}
	var ok bool
	e, ok = r.rm.mapIDvs品目[id]
	if !ok {
		err = xerrors.Errorf("品目が見つかりません。品目ID=%s: %w", id, objects.ErrNotFound)
		return
	}
	return
}

func (r *repTrn品目) GetBy(コード objects.Code品目) (e *objects.Ent品目, err error) {
	if len(r.rm.mapIDvs品目) == 0 {
		err = r.init()
		if err != nil {
			xerrors.Errorf(" :%w", err)
			return
		}
	}
	var ok bool
	e, ok = r.rm.mapコードvs品目[コード]
	if !ok {
		err = xerrors.Errorf("品目が見つかりません。品目コード=%s: %w", コード, objects.ErrNotFound)
		return
	}
	return
}

func (r *repTrn品目) AddNew(e *objects.Ent品目) error {
	// エンティティの責務ではなく、コレクション重複チェックはリポジトリーの責務とする
	if _, err := r.GetBy(e.Getコード); !errors.Is(err, objects.ErrNotFound) {
		return xerrors.Errorf("品目コードがすでに存在します。品目コード=%s: %w", e.Getコード, objects.ErrAlreadyExists)
	}

	r.rm.list品目 = append(r.rm.list品目, e)

	return nil
}

// TODO リソースの一括アップロードやイベント系で1行ずつロギングしない場合はMultiInsertを使う
func (r *repTrn品目) Save(アップロード履歴ID objects.No) (err error) {
	dao品目 := r.rm.dm.NewDaoTrn品目()
	dao仕入品 := r.rm.dm.NewDaoTrn品目仕入品()
	dao製造品 := r.rm.dm.NewDaoTrn品目製造品()
	defer dao品目.Reset()
	defer dao仕入品.Reset()
	defer dao製造品.Reset()

	logger := newCmdTrnリソース変更履歴(r.rm.dm)

	dao単位 := r.rm.dm.NewDaoTrn単位()
	dao生産用品目区分 := r.rm.dm.NewDaoTrn生産用品目区分()

	for _, e := range r.rm.list品目 {
		// インスタンス生成時とインスタンス更新時にチェック済み（になる想定）なのでエラーチェックしない
		dr単位, _ := dao単位.GetByCode(e.Get基準単位.Getコード)
		dr生産用品目区分, _ := dao生産用品目区分.GetByCode(e.Get生産用品目区分.Getコード)

		// 品目を更新 or 追加
		if dr品目, errCode := dao品目.GetByCode(e.Getコード); !errors.Is(errCode, dao.NotFoundError) {
			// 品目
			dr品目.Import(e.Getコード, e.Get名称, dr単位.FldID, dr生産用品目区分.FldID)
			_, err = dao品目.UpdateBy(dr品目)
			if err != nil {
				err = xerrors.Errorf(": %w", err)
				return
			}
			logger.Write(dr品目, アップロード履歴ID)

			// 仕入品
			if dr仕入品, errCode := dao仕入品.GetBy(dr品目.FldID); !errors.Is(errCode, dao.NotFoundError) {
				// すでにあるdr仕入品を更新 or 削除
				if e.Get仕入品 != nil {
					dr仕入品.Import(e.Get仕入品.Get標準単価.Amt(), e.Get仕入品.Get標準単価.Cur(), e.Get仕入品.Get標準単価.PerUnit())
					_, err = dao仕入品.UpdateBy(dr仕入品)
					if err != nil {
						err = xerrors.Errorf(": %w", err)
						return
					}
					logger.Write(dr仕入品, アップロード履歴ID)
				} else {
					_, err = dao仕入品.DeleteBy(dr仕入品)
					if err != nil {
						err = xerrors.Errorf(": %w", err)
						return
					}
					logger.Write(dr品目, アップロード履歴ID)
				}
			} else {
				// dr仕入品を追加
				if e.Get仕入品 != nil {
					dr仕入品 := &dao.Dto品目仕入品{
						FldID:       dr品目.FldID,
						Fld標準単価:     e.Get仕入品.Get標準単価.Amt(),
						Fld標準単価通貨ID: e.Get仕入品.Get標準単価.Cur(),
						Fld標準単価単位ID: e.Get仕入品.Get標準単価.PerUnit(),
						Ub:          dao.NewUb品目仕入品(),
					}
					err = dao仕入品.Insert(dr仕入品)
					if err != nil {
						err = xerrors.Errorf(": %w", err)
						return
					}
					logger.Write(dr仕入品, アップロード履歴ID)
				}
			}

			// 製造品
			if dr製造品, errCode := dao製造品.GetBy(dr品目.FldID); !errors.Is(errCode, dao.NotFoundError) {
				// すでにあるdr製造品を更新 or 削除
				if e.Get製造品 != nil {
					dr製造品.Import(e.Get製造品.GetMRP計算対象フラグ)
					_, err = dao製造品.UpdateBy(dr製造品)
					if err != nil {
						err = xerrors.Errorf(": %w", err)
						return
					}
					logger.Write(dr製造品, アップロード履歴ID)
				} else {
					_, err = dao製造品.DeleteBy(dr製造品)
					if err != nil {
						err = xerrors.Errorf(": %w", err)
						return
					}
					logger.Write(dr製造品, アップロード履歴ID)
				}
			} else {
				// dr製造品を追加
				if e.Get製造品 != nil {
					dr製造品 := &dao.Dto品目製造品{
						FldID:         dr品目.FldID,
						FldMRP計算対象フラグ: e.Get製造品.GetMRP計算対象フラグ,
						Ub:            dao.NewUb品目製造品(),
					}
					err = dao製造品.Insert(dr製造品)
					if err != nil {
						err = xerrors.Errorf(": %w", err)
						return
					}
					logger.Write(dr製造品, アップロード履歴ID)
				}
			}

		} else {
			// 品目を追加
			dr品目 := &dao.Dto品目{
				Fldコード:       e.Getコード,
				Fld名称:        e.Get名称,
				Fld基準単位ID:    dr単位.FldID,
				Fld生産用品目区分ID: dr生産用品目区分.FldID,
				Ub:           dao.NewUb品目(),
			}
			dr品目.FldID, err = dao品目.Insert(dr品目)
			if err != nil {
				err = xerrors.Errorf(": %w", err)
				return
			}
			logger.Write(dr品目, アップロード履歴ID)

			// 仕入品を追加
			if e.Get仕入品 != nil {
				dr仕入品 := &dao.Dto品目仕入品{
					FldID:       dr品目.FldID,
					Fld標準単価:     e.Get仕入品.Get標準単価.Amt(),
					Fld標準単価通貨ID: e.Get仕入品.Get標準単価.Cur(),
					Fld標準単価単位ID: e.Get仕入品.Get標準単価.PerUnit(),
					Ub:          dao.NewUb品目仕入品(),
				}
				err = dao仕入品.Insert(dr仕入品)
				if err != nil {
					err = xerrors.Errorf(": %w", err)
					return
				}
				logger.Write(dr仕入品, アップロード履歴ID)
			}
			// 製造品を追加
			if e.Get製造品 != nil {
				dr製造品 := &dao.Dto品目製造品{
					FldID:         dr品目.FldID,
					FldMRP計算対象フラグ: e.Get製造品.GetMRP計算対象フラグ,
					Ub:            dao.NewUb品目製造品(),
				}
				err = dao製造品.Insert(dr製造品)
				if err != nil {
					err = xerrors.Errorf(": %w", err)
					return
				}
				logger.Write(dr製造品, アップロード履歴ID)
			}
		}
	}

	return
}
