// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"fmt"
	pq "github.com/lib/pq"
)

type Wb品目製造品 interface {
	And(field fld品目製造品, op whereBuilderOperater, val interface{}) Wb品目製造品
	Clear()
	build(argCntStart ...int) (where Where)
}
type wb品目製造品 struct {
	config []whereBuilderExp
}

func NewWb品目製造品() Wb品目製造品 {
	return &wb品目製造品{config: make([]whereBuilderExp, 0)}
}
func newWb品目製造品WithPrimaryKeys(ID Id) Wb品目製造品 {
	wb := &wb品目製造品{config: make([]whereBuilderExp, 0)}
	wb.And(Tbl品目製造品().FldID(), OpEqu, ID)

	return wb
}
func (wb *wb品目製造品) And(field fld品目製造品, op whereBuilderOperater, val interface{}) Wb品目製造品 {
	wb.config = append(wb.config, whereBuilderExp{
		field: string(field),
		op:    op,
		val:   val,
	})
	return wb
}
func (wb *wb品目製造品) Clear() {
	wb.config = make([]whereBuilderExp, 0)
}
func (wb *wb品目製造品) build(argCntStart ...int) (where Where) {
	where.w = ""
	where.prms = make([]interface{}, 0, len(wb.config))
	argCnt := 1
	if len(argCntStart) == 1 {
		argCnt = argCntStart[0]
	}
	for _, e := range wb.config {
		switch e.op {
		case OpIn:
			where.w += fmt.Sprintf(" AND (\"%s\"%s)", e.field, fmt.Sprintf(e.op.string(), fmt.Sprintf("$%d", argCnt)))
			argCnt++
			where.prms = append(where.prms, pq.Array(e.val))
			continue
		case OpIsNull:
			fallthrough
		case OpIsNotNull:
			where.w += fmt.Sprintf(" AND (\"%s\"%s)", e.field, e.op.string())
			continue
		default:
			where.w += fmt.Sprintf(" AND (\"%s\"%s)", e.field, fmt.Sprintf(e.op.string(), fmt.Sprintf("$%d", argCnt)))
			argCnt++
			where.prms = append(where.prms, e.val)
			continue
		}
	}
	return
}

type nothingWb品目製造品 struct{}

func (wb *nothingWb品目製造品) And(field fld品目製造品, op whereBuilderOperater, val interface{}) Wb品目製造品 {
	return wb
}
func (wb *nothingWb品目製造品) Clear() {}
func (wb *nothingWb品目製造品) build(argCntStart ...int) (where Where) {
	return Where{w: " AND 1<>1"}
}
func NewNothingWb品目製造品() Wb品目製造品 {
	return &nothingWb品目製造品{}
}
