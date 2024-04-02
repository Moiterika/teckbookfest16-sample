package infra

import (
	"database/sql"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra/dao"
)

type repManagerDb struct {
	dm *dao.DaoDbManager

	list単位     []*objects.Ent単位
	mapIDvs単位  map[types.Unit]*objects.Ent単位
	mapコードvs単位 map[objects.Code単位]*objects.Ent単位

	list生産用品目区分     []*objects.Ent生産用品目区分
	mapIDvs生産用品目区分  map[dao.Id]*objects.Ent生産用品目区分
	mapコードvs生産用品目区分 map[objects.Code生産用品目区分]*objects.Ent生産用品目区分

	list品目     []*objects.Ent品目
	mapIDvs品目  map[dao.Id]*objects.Ent品目
	mapコードvs品目 map[objects.Code品目]*objects.Ent品目

	list受払    []*objects.Ent受払
	mapIDvs受払 map[objects.No]*objects.Ent受払
}

func NewRepManagerWithDb(db *sql.DB) *repManagerDb {
	return &repManagerDb{
		dm:              dao.NewDaoDbManager(db),
		list単位:          make([]*objects.Ent単位, 0),
		mapIDvs単位:       make(map[types.Unit]*objects.Ent単位),
		mapコードvs単位:      make(map[objects.Code単位]*objects.Ent単位),
		list生産用品目区分:     make([]*objects.Ent生産用品目区分, 0),
		mapIDvs生産用品目区分:  make(map[dao.Id]*objects.Ent生産用品目区分),
		mapコードvs生産用品目区分: make(map[objects.Code生産用品目区分]*objects.Ent生産用品目区分),
		list品目:          make([]*objects.Ent品目, 0),
		mapIDvs品目:       make(map[dao.Id]*objects.Ent品目),
		mapコードvs品目:      make(map[objects.Code品目]*objects.Ent品目),
		list受払:          make([]*objects.Ent受払, 0),
		mapIDvs受払:       make(map[objects.No]*objects.Ent受払),
	}
}

func (rm *repManagerDb) rep単位() *repDb単位 {
	return &repDb単位{
		rm: rm,
	}
}

func (rm *repManagerDb) NewRep単位() objects.Rep単位 {
	return &repDb単位{
		rm: rm,
	}
}

func (rm *repManagerDb) rep生産用品目区分() *repDb生産用品目区分 {
	return &repDb生産用品目区分{
		rm: rm,
	}
}

func (rm *repManagerDb) NewRep生産用品目区分() objects.Rep生産用品目区分 {
	return &repDb生産用品目区分{
		rm: rm,
	}
}

func (rm *repManagerDb) rep品目() *repDb品目 {
	return &repDb品目{
		rm: rm,
	}
}

func (rm *repManagerDb) NewRep品目() objects.Rep品目 {
	return &repDb品目{
		rm: rm,
	}
}

func (rm *repManagerDb) NewRep受払() objects.Rep受払 {
	return &repDb受払{
		rm: rm,
	}
}
