// Code generated by xlsx2go.go; DO NOT EDIT.

package objects

import types "techbookfest16-sample/domain/types"

type Rep品目 interface {
	List() ([]*Ent品目, error)
	GetBy(types.Code品目) (*Ent品目, error)
	AddNew(*Ent品目) error
	Save(types.No) error
}
