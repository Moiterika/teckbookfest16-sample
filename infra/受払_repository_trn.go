package infra

import (
	"errors"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra/dao"

	"github.com/Moiterika/a"

	"golang.org/x/xerrors"
)

type repTrn受払 struct {
	rm *repManagerTrn
}

func (r *repTrn受払) init() (err error) {
	r.rm.list受払, err = r.list()
	if err != nil {
		xerrors.Errorf(" :%w", err)
		return
	}
	r.rm.mapIDvs受払 = a.ToMap(r.rm.list受払, func(e *objects.Ent受払) types.No {
		return e.GetNo
	})

	return
}

func (r *repTrn受払) list() (list []*objects.Ent受払, err error) {
	dt受払, errDt := r.rm.dm.NewDaoTrn受払().Dt()
	if errDt != nil {
		err = xerrors.Errorf(" :%w", errDt)
		return
	}

	rep品目 := r.rm.rep品目()
	dao仕入 := r.rm.dm.NewDaoTrn受払仕入()
	dao単位 := r.rm.dm.NewDaoTrn単位()

	list = make([]*objects.Ent受払, len(dt受払))
	for i, dr := range dt受払 {
		品目, err品目 := rep品目.getBy(dr.Fld品目ID)
		if err品目 != nil {
			err = xerrors.Errorf(" :%w", err品目)
			return
		}
		dr仕入, err仕入 := dao仕入.GetBy(dr.FldNo)
		if err仕入 != nil && !errors.Is(err仕入, dao.NotFoundError) {
			err = xerrors.Errorf(" :%w", err仕入)
			return
		}
		dr仕入単位, _ := dao単位.GetBy(dr仕入.Fld仕入単位ID)
		仕入数量, err仕入数量 := types.NewQuantity(dr仕入.Fld仕入数量, dr仕入単位.Fldコード)
		if err仕入数量 != nil {
			err = xerrors.Errorf(" :%w", err仕入数量)
			return
		}
		受払仕入, err受払仕入 := objects.NewVal受払仕入(仕入数量,
			types.NewAmount(dr仕入.Fld仕入金額, dr仕入.Fld仕入通貨ID),
			types.NewPrice(dr仕入.Fld仕入単価, dr仕入.Fld仕入通貨ID, dr仕入単位.Fldコード),
		)
		if err受払仕入 != nil {
			err = xerrors.Errorf(" :%w", err受払仕入)
			return
		}
		dr単位, _ := dao単位.GetBy(dr.Fld基準単位ID)
		e := &objects.Ent受払{
			GetNo:    types.No(dr.FldNo),
			Get登録日時:  dr.Fld登録日時,
			Get計上月:   dr.Fld計上月,
			Get受払区分:  dr.Fld受払区分,
			Get赤伝フラグ: dr.Fld赤伝フラグ,
			Get品目:    品目,
			Get基準数量:  types.NewInventory(dr.Fld基準数量, dr単位.Fldコード),
			Get仕入:    受払仕入,
			// 今回は全部nil
			// Get出荷:    &objects.Val受払出荷{},
			// Get投入実績:  &objects.Val受払投入実績{},
			// Get製造実績:  &objects.Val受払製造実績{},
		}
		list[i] = e
	}
	return
}

func (r *repTrn受払) List() ([]*objects.Ent受払, error) {
	if len(r.rm.list受払) == 0 {
		err := r.init()
		if err != nil {
			xerrors.Errorf(" :%w", err)
			return nil, err
		}
	}
	return r.rm.list受払, nil
}

func (r *repTrn受払) GetBy(no types.No) (e *objects.Ent受払, err error) {
	if len(r.rm.mapIDvs受払) == 0 {
		err = r.init()
		if err != nil {
			xerrors.Errorf(" :%w", err)
			return
		}
	}
	var ok bool
	e, ok = r.rm.mapIDvs受払[no]
	if !ok {
		err = xerrors.Errorf("受払が見つかりません。受払ID=%d: %w", no, objects.ErrNotFound)
		return
	}
	return
}

func (r *repTrn受払) AddNew(e *objects.Ent受払) error {
	r.rm.list受払 = append(r.rm.list受払, e)

	return nil
}

// TODO リソースの一括アップロードやイベント系で1行ずつロギングしない場合はMultiInsertを使う
func (r *repTrn受払) Save(アップロード履歴ID types.No) (err error) {
	dao受払 := r.rm.dm.NewDaoTrn受払()
	dao仕入 := r.rm.dm.NewDaoTrn受払仕入()
	defer dao受払.Reset()
	defer dao仕入.Reset()

	dao品目 := r.rm.dm.NewDaoTrn品目()
	dao単位 := r.rm.dm.NewDaoTrn単位()

	for _, e := range r.rm.list受払 {
		dr品目, err品目 := dao品目.GetByCode(e.Get品目.Getコード)
		if err品目 != nil {
			err = xerrors.Errorf(" :%w", err品目)
			return
		}
		dr単位, _ := dao単位.GetByCode(e.Get基準数量.Unit())

		// イベントを更新してよいかは要検討
		if dr, errCode := dao受払.GetBy(dao.Id(e.GetNo)); !errors.Is(errCode, dao.NotFoundError) {
			dr.Import(e.Get登録日時, e.Get計上月, e.Get受払区分, e.Get赤伝フラグ, dr品目.FldID, e.Get基準数量.Val(), dr単位.FldID)
			_, err = dao受払.UpdateBy(dr)
			if err != nil {
				err = xerrors.Errorf(": %w", err)
				return
			}
		} else {
			dr := &dao.Dto受払{
				// FldNo:     0,
				Fld登録日時:   e.Get登録日時,
				Fld計上月:    e.Get計上月,
				Fld受払区分:   e.Get受払区分,
				Fld赤伝フラグ:  e.Get赤伝フラグ,
				Fld品目ID:   dr品目.FldID,
				Fld基準数量:   e.Get基準数量.Val(),
				Fld基準単位ID: dr単位.FldID,
				Ub:        dao.NewUb受払(),
			}
			dr.FldNo, err = dao受払.Insert(dr)
			if err != nil {
				err = xerrors.Errorf(": %w", err)
				return
			}
			e.GetNo = types.No(dr.FldNo)

			if e.Get仕入 != nil {
				dr仕入単位, _ := dao単位.GetByCode(e.Get仕入.Get仕入数量.Unit())
				dr仕入 := &dao.Dto受払仕入{
					FldNo:     dr.FldNo,
					Fld仕入数量:   e.Get仕入.Get仕入数量.Val(),
					Fld仕入単位ID: dr仕入単位.FldID,
					Fld仕入金額:   e.Get仕入.Get仕入金額.Val(),
					Fld仕入通貨ID: e.Get仕入.Get仕入金額.Unit(),
					Fld仕入単価:   e.Get仕入.Get仕入単価.Amt(),
					Ub:        dao.NewUb受払仕入(),
				}
				err = dao仕入.Insert(dr仕入)
				if err != nil {
					err = xerrors.Errorf(": %w", err)
					return
				}
			}

			// 今回は仕入以外、全部nil
		}
	}

	return
}
