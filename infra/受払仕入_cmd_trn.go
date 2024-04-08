package infra

import (
	"fmt"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra/dao"
	"time"
)

type cmdTrn受払仕入 struct {
	rm *repManagerTrn
}

func (c *cmdTrn受払仕入) Entry(
	登録日時 time.Time,
	計上月 time.Time,
	受払区分 objects.Enum受払区分,
	赤伝フラグ bool,
	品目 *objects.Ent品目,
	基準数量 types.Quantity,
	仕入数量 types.Quantity,
	仕入金額 types.Amount,
	仕入単価 types.Price,
) error {
	dao受払 := c.rm.dm.NewDaoTrn受払()
	dao品目 := c.rm.dm.NewDaoTrn品目()
	dao単位 := c.rm.dm.NewDaoTrn単位()

	dr品目, err := dao品目.GetByCode(品目.Getコード)
	if err != nil {
		return fmt.Errorf(" :%w", err)
	}

	if dr基準単位, err := dao単位.GetBy(dr品目.Fld基準単位ID); err != nil {
		return fmt.Errorf(" :%w", err)
	} else if 基準数量.Unit() != dr基準単位.Fldコード {
		return fmt.Errorf("基準単位コードは%sであるべきです。基準単位コード=%s :%w", dr基準単位.Fldコード, 基準数量.Unit(), err)
	}

	dr仕入単位, err := dao単位.GetByCode(仕入数量.Unit())
	if err != nil {
		return fmt.Errorf(" :%w", err)
	}

	if 仕入数量.Unit() != 仕入単価.PerUnit() {
		return fmt.Errorf("仕入数量と仕入単価の数量単位が異なります。 :%w", err)
	}

	dr受払 := &dao.Dto受払{
		//FldNo:     0,
		Fld登録日時:   登録日時,
		Fld計上月:    計上月,
		Fld受払区分:   受払区分,
		Fld赤伝フラグ:  赤伝フラグ,
		Fld品目ID:   dr品目.FldID,
		Fld基準数量:   基準数量.Val(),
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
		Fld仕入数量:   仕入数量.Val(),
		Fld仕入単位ID: dr仕入単位.FldID,
		Fld仕入金額:   仕入金額.Val(),
		Fld仕入通貨ID: 仕入金額.Unit(),
		Fld仕入単価:   仕入単価.Amt(),
		Ub:        dao.NewUb受払仕入(),
	}
	err = dao仕入.Insert(dr仕入)
	if err != nil {
		return err
	}

	return nil
}
