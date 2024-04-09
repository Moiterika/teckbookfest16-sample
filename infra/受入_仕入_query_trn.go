package infra

import (
	"errors"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"

	"golang.org/x/xerrors"
)

type qryTrn受払仕入 struct {
	rm       *repManagerTrn
	isLoaded bool
}

func NewQryTrn受払仕入(rm *repManagerTrn) objects.Qry受払仕入 {
	return &qryTrn受払仕入{
		rm: rm,
	}
}

func (r *qryTrn受払仕入) init() error {
	q := NewQryTrn受払(r.rm)
	rep単位 := r.rm.rep単位()

	// TODO 計上年月で絞るなどしないと、全件取得はそのうち破綻する。
	dao仕入 := r.rm.dm.NewDaoTrn受払仕入()
	dt仕入, err := dao仕入.Dt()
	if err != nil {
		return xerrors.Errorf(" :%w", err)
	}
	for i, dr := range dt仕入 {
		e単位, err := rep単位.getBy(dr.Fld仕入単位ID)
		if err != nil && errors.Is(err, objects.ErrNotFound) {
			return xerrors.Errorf(" :%w", err)
		}
		e受払, notFound := q.GetBy(types.No(dr.FldNo))
		if notFound != nil {
			return xerrors.Errorf("受払No=%d :%w", dr.FldNo, notFound)
		}
		仕入数量, err := types.NewQuantity(dr.Fld仕入数量, e単位.Getコード)
		if err != nil {
			return xerrors.Errorf(" :%w", err)
		}
		e, err := objects.NewEnt受払仕入(e受払,
			仕入数量,
			types.NewAmount(dr.Fld仕入金額, dr.Fld仕入通貨ID),
			types.NewPrice(dr.Fld仕入単価, dr.Fld仕入通貨ID, e単位.Getコード),
		)
		if err != nil {
			return xerrors.Errorf(" :%w", err)
		}
		r.rm.list仕入[i] = e
		r.rm.mapNovs仕入[e.GetNo] = e
	}

	r.isLoaded = true

	return nil
}

func (r *qryTrn受払仕入) List() ([]*objects.Ent受払仕入, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	return r.rm.list仕入, nil
}

func (r *qryTrn受払仕入) GetBy(no types.No) (*objects.Ent受払仕入, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapNovs仕入[no]
	if !ok {
		return nil, xerrors.Errorf("仕入が見つかりません。受払No=%d: %w", no, objects.ErrNotFound)
	}
	return e, nil
}
