// Code generated by xlsx2go.go; DO NOT EDIT.

package objects

import types "techbookfest16-sample/domain/types"

type Qry受払仕入 interface {
	List() ([]*Ent受払仕入, error)
	GetBy(types.No) (*Ent受払仕入, error)
}
