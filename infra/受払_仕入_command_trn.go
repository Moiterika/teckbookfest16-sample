package infra

import (
	"fmt"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/infra/dao"
)

type cmdTrn受払仕入 struct {
	rm *repManagerTrn
}

func NewCmdTrn受払(rm *repManagerTrn) objects.Cmd受払仕入 {
	return &cmdTrn受払仕入{
		rm: rm,
	}
}

func (c *cmdTrn受払仕入) Entry(
	e *objects.Ent受払仕入,
) error {
	dao受払 := c.rm.dm.NewDaoTrn受払()
	dao品目 := c.rm.dm.NewDaoTrn品目()
	dao単位 := c.rm.dm.NewDaoTrn単位()

	dr品目, err := dao品目.GetByCode(e.Get品目.Getコード)
	if err != nil {
		return fmt.Errorf(" :%w", err)
	}

	if dr基準単位, err := dao単位.GetBy(dr品目.Fld基準単位ID); err != nil {
		return fmt.Errorf(" :%w", err)
	} else if e.Get基準数量.Unit() != dr基準単位.Fldコード {
		return fmt.Errorf("基準単位コードは%sであるべきです。基準単位コード=%s :%w", dr基準単位.Fldコード, e.Get基準数量.Unit(), err)
	}

	dr仕入単位, err := dao単位.GetByCode(e.Get仕入数量.Unit())
	if err != nil {
		return fmt.Errorf(" :%w", err)
	}

	if e.Get仕入数量.Unit() != e.Get仕入単価.PerUnit() {
		return fmt.Errorf("仕入数量と仕入単価の数量単位が異なります。 :%w", err)
	}

	dr受払 := &dao.Dto受払{
		//FldNo:     0,
		Fld登録日時:   e.Get登録日時,
		Fld計上月:    e.Get計上月,
		Fld受払区分:   e.Get受払区分,
		Fld赤伝フラグ:  e.Get赤伝フラグ,
		Fld品目ID:   dr品目.FldID,
		Fld基準数量:   e.Get基準数量.Val(),
		Fld基準単位ID: dr品目.Fld基準単位ID,
		Ub:        dao.NewUb受払(),
	}
	err = dao受払.Insert(dr受払)
	if err != nil {
		return err
	}

	dao仕入 := c.rm.dm.NewDaoTrn受払仕入()
	dr仕入 := &dao.Dto受払仕入{
		FldNo:     dr受払.FldNo,
		Fld仕入数量:   e.Get仕入数量.Val(),
		Fld仕入単位ID: dr仕入単位.FldID,
		Fld仕入金額:   e.Get仕入金額.Val(),
		Fld仕入通貨ID: e.Get仕入金額.Unit(),
		Fld仕入単価:   e.Get仕入単価.Amt(),
		Ub:        dao.NewUb受払仕入(),
	}
	err = dao仕入.Insert(dr仕入)
	if err != nil {
		return err
	}

	return nil
}
