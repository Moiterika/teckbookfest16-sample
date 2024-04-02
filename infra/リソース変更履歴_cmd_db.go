package infra

import (
	"log"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/infra/dao"
	"time"
)

type cmdDbリソース変更履歴 struct {
	dm *dao.DaoDbManager
}

func newCmdDbリソース変更履歴(dm *dao.DaoDbManager) *cmdDbリソース変更履歴 {
	return &cmdDbリソース変更履歴{
		dm: dm,
	}
}

func (c *cmdDbリソース変更履歴) Write(dr ResourceDataRow, アップロード履歴ID objects.No) {
	// TODO DBのバージョン情報も記録しておくか
	// daoMig := c.dm.NewDaoTrngorpmigrations()
	// daoMig.WbForInit.And(dao.Tblgorpmigrations().Fldapplied_at(), dao.OpIsNotNull, nil)
	// mig, err := daoMig.Dt()
	// if err != nil {
	// 	log.Printf("%v", err)
	// 	return
	// }
	// dbVersion := a.LastOrDefault(mig).Fldapplied_at.Time

	if dr.RowState() == dao.UnChanged {
		return
	}

	b, err := dr.ToJson()
	if err != nil {
		log.Printf("%v", err)
		return
	}

	リソース変更履歴 := &dao.Dtoリソース変更履歴{
		Fld登録日時:  time.Now(),
		Fldリソース名: dr.TableName(),
		Fld変更区分:  dr.RowState().String(),
		Fld変更内容:  b,
		//Ub:       dao.NewUbリソース変更履歴(),
	}
	id, err := c.dm.NewDaoDbリソース変更履歴().Insert(リソース変更履歴)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	リソース変更履歴アップロード時 := &dao.Dtoリソース変更履歴アップロード時{
		FldID:         id,
		Fldアップロード履歴ID: dao.Id(アップロード履歴ID),
		//Ub:            dao.NewUbリソース変更履歴アップロード時(),
	}
	err = c.dm.NewDaoDbリソース変更履歴アップロード時().Insert(リソース変更履歴アップロード時)
	if err != nil {
		log.Printf("%v", err)
		return
	}
}
