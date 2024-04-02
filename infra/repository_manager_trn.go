package infra

import (
	"database/sql"
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"techbookfest16-sample/infra/dao"
)

type repManagerTrn struct {
	dm *dao.DaoTrnManager

	list単位     []*objects.Ent単位
	mapIDvs単位  map[dao.Id]*objects.Ent単位
	mapコードvs単位 map[types.Code単位]*objects.Ent単位

	list生産用品目区分     []*objects.Ent生産用品目区分
	mapIDvs生産用品目区分  map[dao.Id]*objects.Ent生産用品目区分
	mapコードvs生産用品目区分 map[types.Code生産用品目区分]*objects.Ent生産用品目区分

	list品目     []*objects.Ent品目
	mapIDvs品目  map[dao.Id]*objects.Ent品目
	mapコードvs品目 map[types.Code品目]*objects.Ent品目

	list受払    []*objects.Ent受払
	mapIDvs受払 map[objects.No]*objects.Ent受払
}

func NewRepManagerWithTrn(trn *sql.Tx) *repManagerTrn {
	return &repManagerTrn{
		dm:              dao.NewDaoTrnManager(trn),
		list単位:          make([]*objects.Ent単位, 0),
		mapIDvs単位:       make(map[dao.Id]*objects.Ent単位),
		mapコードvs単位:      make(map[types.Code単位]*objects.Ent単位),
		list生産用品目区分:     make([]*objects.Ent生産用品目区分, 0),
		mapIDvs生産用品目区分:  make(map[dao.Id]*objects.Ent生産用品目区分),
		mapコードvs生産用品目区分: make(map[types.Code生産用品目区分]*objects.Ent生産用品目区分),
		list品目:          make([]*objects.Ent品目, 0),
		mapIDvs品目:       make(map[dao.Id]*objects.Ent品目),
		mapコードvs品目:      make(map[types.Code品目]*objects.Ent品目),
		list受払:          make([]*objects.Ent受払, 0),
		mapIDvs受払:       make(map[objects.No]*objects.Ent受払),
	}
}

func (rm *repManagerTrn) rep単位() *repTrn単位 {
	return &repTrn単位{
		rm: rm,
	}
}

func (rm *repManagerTrn) NewRep単位() objects.Rep単位 {
	return &repTrn単位{
		rm: rm,
	}
}

func (rm *repManagerTrn) rep生産用品目区分() *repTrn生産用品目区分 {
	return &repTrn生産用品目区分{
		rm: rm,
	}
}

func (rm *repManagerTrn) NewRep生産用品目区分() objects.Rep生産用品目区分 {
	return &repTrn生産用品目区分{
		rm: rm,
	}
}

func (rm *repManagerTrn) rep品目() *repTrn品目 {
	return &repTrn品目{
		rm: rm,
	}
}

func (rm *repManagerTrn) NewRep品目() objects.Rep品目 {
	return &repTrn品目{
		rm: rm,
	}
}

func (rm *repManagerTrn) NewRep受払() objects.Rep受払 {
	return &repTrn受払{
		rm: rm,
	}
}
