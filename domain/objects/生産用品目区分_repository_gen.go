// Code generated by xlsx2go.go; DO NOT EDIT.

package objects

import types "techbookfest16-sample/domain/types"

type Rep生産用品目区分 interface {
	List() ([]*Ent生産用品目区分, error)
	GetBy(types.Code生産用品目区分) (*Ent生産用品目区分, error)
	AddNew(*Ent生産用品目区分) error
	Save(types.No) error
}
