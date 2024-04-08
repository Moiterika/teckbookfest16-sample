package infra

import (
	"errors"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"

	"golang.org/x/xerrors"
)

type qryTrn受払 struct {
	rm       *repManagerTrn
	isLoaded bool
}

func (r *qryTrn受払) init() error {
	// TODO 計上年月で絞るなどしないと、全件取得はそのうち破綻する。
	// サンプルなのでいったんヨシ
	dt受払, err := r.rm.dm.NewDaoTrn受払().Dt()
	if err != nil {
		return xerrors.Errorf(" :%w", err)
	}

	rep品目 := r.rm.rep品目()
	rep単位 := r.rm.rep単位()

	r.rm.list受払 = make([]*objects.Ent受払, len(dt受払))
	for i, dr := range dt受払 {
		品目, notFound := rep品目.getBy(dr.Fld品目ID)
		if notFound != nil {
			return xerrors.Errorf(" :%w", notFound)
		}
		e単位, notFound := rep単位.getBy(dr.Fld基準単位ID)
		if notFound != nil {
			return xerrors.Errorf(" :%w", notFound)
		}
		e, notFound := objects.NewEnt受払(
			types.No(dr.FldNo),
			dr.Fld登録日時,
			dr.Fld計上月,
			dr.Fld受払区分,
			dr.Fld赤伝フラグ,
			品目,
			types.NewInventory(dr.Fld基準数量, e単位.Getコード),
		)
		if notFound != nil {
			return xerrors.Errorf(" :%w", notFound)
		}
		r.rm.list受払[i] = e
		r.rm.mapNovs受払[e.GetNo] = e
	}

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
		e受払, ok := r.rm.mapNovs受払[types.No(dr.FldNo)]
		if !ok {
			return xerrors.Errorf("受払No=%d :%w", dr.FldNo, objects.ErrNotFound)
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

func (r *qryTrn受払) List() ([]*objects.Ent受払, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	return r.rm.list受払, nil
}

func (r *qryTrn受払) List仕入() ([]*objects.Ent受払仕入, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	return r.rm.list仕入, nil
}

func (r *qryTrn受払) GetBy(no types.No) (*objects.Ent受払, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapNovs受払[no]
	if !ok {
		return nil, xerrors.Errorf("受払が見つかりません。受払ID=%d: %w", no, objects.ErrNotFound)
	}
	return e, nil
}
