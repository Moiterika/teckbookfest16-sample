package infra

import (
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"

	"golang.org/x/xerrors"
)

type qryTrn受払 struct {
	rm       *repManagerTrn
	isLoaded bool
}

func NewQryTrn受払(rm *repManagerTrn) objects.Qry受払 {
	return &qryTrn受払{
		rm: rm,
	}
}

func (r *qryTrn受払) init() error {
	// TODO 計上年月で絞るなどしないと、全件取得はそのうち破綻する。
	// サンプルなのでいったんヨシ
	dt受払, err := r.rm.dm.NewDaoTrn受払With(r.rm.repManagerArgs.wb受払).Dt()
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
		e, err := objects.NewEnt受払(
			dr.Fld登録日時,
			dr.Fld計上月,
			dr.Fld受払区分,
			dr.Fld赤伝フラグ,
			品目,
			types.NewInventory(dr.Fld基準数量, e単位.Getコード),
		)
		if err != nil {
			return xerrors.Errorf(" :%w", err)
		}
		e.GetNo = types.No(dr.FldNo)

		r.rm.list受払[i] = e
		r.rm.mapNovs受払[e.GetNo] = e
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

func (r *qryTrn受払) GetBy(no types.No) (*objects.Ent受払, error) {
	if !r.isLoaded {
		err := r.init()
		if err != nil {
			return nil, xerrors.Errorf(" :%w", err)
		}
	}
	e, ok := r.rm.mapNovs受払[no]
	if !ok {
		return nil, xerrors.Errorf("受払が見つかりません。受払No=%d: %w", no, types.ErrNotFound)
	}
	return e, nil
}
