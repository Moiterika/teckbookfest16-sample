// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"database/sql"
	types "techbookfest16-sample/domain/types"
)

// DaoTrnManager is a struct.
// This struct does not include sync.Mutex lock system.
// The sync.Mutex is YAGNI(You ain't gonna need it). This struct will change, when this struct is simultaneously accessed from some instances; the one of usage example is the singleton pattaern called from anywhere in backend system
type DaoTrnManager struct {
	trn                      *sql.Tx
	dtEnumログ区分               []*DtoEnumログ区分
	dtEnum受払区分               []*DtoEnum受払区分
	dtgorpmigrations         []*Dtogorpmigrations
	dtリソース変更履歴               []*Dtoリソース変更履歴
	mapIDvsDrリソース変更履歴        map[Id]*Dtoリソース変更履歴
	dtリソース変更履歴アップロード時        []*Dtoリソース変更履歴アップロード時
	mapIDvsDrリソース変更履歴アップロード時 map[Id]*Dtoリソース変更履歴アップロード時
	dtログ                     []*Dtoログ
	mapIDvsDrログ              map[Id]*Dtoログ
	dtログアップロード時              []*Dtoログアップロード時
	mapIDvsDrログアップロード時       map[Id]*Dtoログアップロード時
	dtログ画面操作時                []*Dtoログ画面操作時
	mapIDvsDrログ画面操作時         map[Id]*Dtoログ画面操作時
	dt単位                     []*Dto単位
	mapIDvsDr単位              map[Id]*Dto単位
	mapコードvsDr単位             map[types.Code単位]*Dto単位
	dt受払                     []*Dto受払
	mapIDvsDr受払              map[Id]*Dto受払
	dt受払仕入                   []*Dto受払仕入
	mapIDvsDr受払仕入            map[Id]*Dto受払仕入
	dt受払出荷                   []*Dto受払出荷
	mapIDvsDr受払出荷            map[Id]*Dto受払出荷
	dt受払投入実績                 []*Dto受払投入実績
	mapIDvsDr受払投入実績          map[Id]*Dto受払投入実績
	dt受払製造実績                 []*Dto受払製造実績
	mapIDvsDr受払製造実績          map[Id]*Dto受払製造実績
	dt品目                     []*Dto品目
	mapIDvsDr品目              map[Id]*Dto品目
	mapコードvsDr品目             map[types.Code品目]*Dto品目
	dt品目仕入品                  []*Dto品目仕入品
	mapIDvsDr品目仕入品           map[Id]*Dto品目仕入品
	dt品目製造品                  []*Dto品目製造品
	mapIDvsDr品目製造品           map[Id]*Dto品目製造品
	dt生産用品目区分                []*Dto生産用品目区分
	mapIDvsDr生産用品目区分         map[Id]*Dto生産用品目区分
	mapコードvsDr生産用品目区分        map[types.Code生産用品目区分]*Dto生産用品目区分
}

func NewDaoTrnManager(trn *sql.Tx) *DaoTrnManager {
	return &DaoTrnManager{
		mapIDvsDrリソース変更履歴:        make(map[Id]*Dtoリソース変更履歴),
		mapIDvsDrリソース変更履歴アップロード時: make(map[Id]*Dtoリソース変更履歴アップロード時),
		mapIDvsDrログ:        make(map[Id]*Dtoログ),
		mapIDvsDrログアップロード時: make(map[Id]*Dtoログアップロード時),
		mapIDvsDrログ画面操作時:   make(map[Id]*Dtoログ画面操作時),
		mapIDvsDr単位:        make(map[Id]*Dto単位),
		mapIDvsDr受払:        make(map[Id]*Dto受払),
		mapIDvsDr受払仕入:      make(map[Id]*Dto受払仕入),
		mapIDvsDr受払出荷:      make(map[Id]*Dto受払出荷),
		mapIDvsDr受払投入実績:    make(map[Id]*Dto受払投入実績),
		mapIDvsDr受払製造実績:    make(map[Id]*Dto受払製造実績),
		mapIDvsDr品目:        make(map[Id]*Dto品目),
		mapIDvsDr品目仕入品:     make(map[Id]*Dto品目仕入品),
		mapIDvsDr品目製造品:     make(map[Id]*Dto品目製造品),
		mapIDvsDr生産用品目区分:   make(map[Id]*Dto生産用品目区分),
		mapコードvsDr単位:       make(map[types.Code単位]*Dto単位),
		mapコードvsDr品目:       make(map[types.Code品目]*Dto品目),
		mapコードvsDr生産用品目区分:  make(map[types.Code生産用品目区分]*Dto生産用品目区分),
		trn:                trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrnEnumログ区分() daoTrnEnumログ区分 {
	return daoTrnEnumログ区分{
		WbForInit: NewWbEnumログ区分(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrnEnumログ区分With(wb WbEnumログ区分) daoTrnEnumログ区分 {
	if wb != nil {
		return daoTrnEnumログ区分{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrnEnumログ区分()
}
func (dm *DaoTrnManager) NewDaoTrnEnum受払区分() daoTrnEnum受払区分 {
	return daoTrnEnum受払区分{
		WbForInit: NewWbEnum受払区分(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrnEnum受払区分With(wb WbEnum受払区分) daoTrnEnum受払区分 {
	if wb != nil {
		return daoTrnEnum受払区分{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrnEnum受払区分()
}
func (dm *DaoTrnManager) NewDaoTrngorpmigrations() daoTrngorpmigrations {
	return daoTrngorpmigrations{
		WbForInit: NewWbgorpmigrations(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrngorpmigrationsWith(wb Wbgorpmigrations) daoTrngorpmigrations {
	if wb != nil {
		return daoTrngorpmigrations{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrngorpmigrations()
}
func (dm *DaoTrnManager) NewDaoTrnリソース変更履歴() daoTrnリソース変更履歴 {
	return daoTrnリソース変更履歴{
		WbForInit: NewWbリソース変更履歴(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrnリソース変更履歴With(wb Wbリソース変更履歴) daoTrnリソース変更履歴 {
	if wb != nil {
		return daoTrnリソース変更履歴{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrnリソース変更履歴()
}
func (dm *DaoTrnManager) NewDaoTrnリソース変更履歴アップロード時() daoTrnリソース変更履歴アップロード時 {
	return daoTrnリソース変更履歴アップロード時{
		WbForInit: NewWbリソース変更履歴アップロード時(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrnリソース変更履歴アップロード時With(wb Wbリソース変更履歴アップロード時) daoTrnリソース変更履歴アップロード時 {
	if wb != nil {
		return daoTrnリソース変更履歴アップロード時{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrnリソース変更履歴アップロード時()
}
func (dm *DaoTrnManager) NewDaoTrnログ() daoTrnログ {
	return daoTrnログ{
		WbForInit: NewWbログ(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrnログWith(wb Wbログ) daoTrnログ {
	if wb != nil {
		return daoTrnログ{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrnログ()
}
func (dm *DaoTrnManager) NewDaoTrnログアップロード時() daoTrnログアップロード時 {
	return daoTrnログアップロード時{
		WbForInit: NewWbログアップロード時(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrnログアップロード時With(wb Wbログアップロード時) daoTrnログアップロード時 {
	if wb != nil {
		return daoTrnログアップロード時{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrnログアップロード時()
}
func (dm *DaoTrnManager) NewDaoTrnログ画面操作時() daoTrnログ画面操作時 {
	return daoTrnログ画面操作時{
		WbForInit: NewWbログ画面操作時(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrnログ画面操作時With(wb Wbログ画面操作時) daoTrnログ画面操作時 {
	if wb != nil {
		return daoTrnログ画面操作時{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrnログ画面操作時()
}
func (dm *DaoTrnManager) NewDaoTrn単位() daoTrn単位 {
	return daoTrn単位{
		WbForInit: NewWb単位(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn単位With(wb Wb単位) daoTrn単位 {
	if wb != nil {
		return daoTrn単位{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn単位()
}
func (dm *DaoTrnManager) NewDaoTrn受払() daoTrn受払 {
	return daoTrn受払{
		WbForInit: NewWb受払(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn受払With(wb Wb受払) daoTrn受払 {
	if wb != nil {
		return daoTrn受払{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn受払()
}
func (dm *DaoTrnManager) NewDaoTrn受払仕入() daoTrn受払仕入 {
	return daoTrn受払仕入{
		WbForInit: NewWb受払仕入(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn受払仕入With(wb Wb受払仕入) daoTrn受払仕入 {
	if wb != nil {
		return daoTrn受払仕入{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn受払仕入()
}
func (dm *DaoTrnManager) NewDaoTrn受払出荷() daoTrn受払出荷 {
	return daoTrn受払出荷{
		WbForInit: NewWb受払出荷(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn受払出荷With(wb Wb受払出荷) daoTrn受払出荷 {
	if wb != nil {
		return daoTrn受払出荷{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn受払出荷()
}
func (dm *DaoTrnManager) NewDaoTrn受払投入実績() daoTrn受払投入実績 {
	return daoTrn受払投入実績{
		WbForInit: NewWb受払投入実績(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn受払投入実績With(wb Wb受払投入実績) daoTrn受払投入実績 {
	if wb != nil {
		return daoTrn受払投入実績{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn受払投入実績()
}
func (dm *DaoTrnManager) NewDaoTrn受払製造実績() daoTrn受払製造実績 {
	return daoTrn受払製造実績{
		WbForInit: NewWb受払製造実績(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn受払製造実績With(wb Wb受払製造実績) daoTrn受払製造実績 {
	if wb != nil {
		return daoTrn受払製造実績{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn受払製造実績()
}
func (dm *DaoTrnManager) NewDaoTrn品目() daoTrn品目 {
	return daoTrn品目{
		WbForInit: NewWb品目(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn品目With(wb Wb品目) daoTrn品目 {
	if wb != nil {
		return daoTrn品目{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn品目()
}
func (dm *DaoTrnManager) NewDaoTrn品目仕入品() daoTrn品目仕入品 {
	return daoTrn品目仕入品{
		WbForInit: NewWb品目仕入品(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn品目仕入品With(wb Wb品目仕入品) daoTrn品目仕入品 {
	if wb != nil {
		return daoTrn品目仕入品{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn品目仕入品()
}
func (dm *DaoTrnManager) NewDaoTrn品目製造品() daoTrn品目製造品 {
	return daoTrn品目製造品{
		WbForInit: NewWb品目製造品(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn品目製造品With(wb Wb品目製造品) daoTrn品目製造品 {
	if wb != nil {
		return daoTrn品目製造品{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn品目製造品()
}
func (dm *DaoTrnManager) NewDaoTrn生産用品目区分() daoTrn生産用品目区分 {
	return daoTrn生産用品目区分{
		WbForInit: NewWb生産用品目区分(),
		dm:        dm,
		trn:       dm.trn,
	}
}
func (dm *DaoTrnManager) NewDaoTrn生産用品目区分With(wb Wb生産用品目区分) daoTrn生産用品目区分 {
	if wb != nil {
		return daoTrn生産用品目区分{
			WbForInit: wb,
			dm:        dm,
			trn:       dm.trn,
		}
	}
	return dm.NewDaoTrn生産用品目区分()
}
