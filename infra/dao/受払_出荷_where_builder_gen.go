// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"fmt"
	pq "github.com/lib/pq"
)

type Wb受払出荷 interface {
	And(field fld受払出荷, op whereBuilderOperater, val interface{}) Wb受払出荷
	Clear()
	Exists(Eb受払出荷) Wb受払出荷
	build(argCntStart ...int) (where Where)
}
type wb受払出荷 struct {
	config []whereBuilderExp
	ebs    []Eb受払出荷
}

func NewWb受払出荷() Wb受払出荷 {
	return &wb受払出荷{
		config: make([]whereBuilderExp, 0),
		ebs:    make([]Eb受払出荷, 0),
	}
}
func newWb受払出荷WithPrimaryKeys(No Id) Wb受払出荷 {
	wb := &wb受払出荷{config: make([]whereBuilderExp, 0)}
	wb.And(Tbl受払出荷().FldNo(), OpEqu, No)

	return wb
}
func (wb *wb受払出荷) And(field fld受払出荷, op whereBuilderOperater, val interface{}) Wb受払出荷 {
	wb.config = append(wb.config, whereBuilderExp{
		field: string(field),
		op:    op,
		val:   val,
	})
	return wb
}
func (wb *wb受払出荷) Clear() {
	wb.config = make([]whereBuilderExp, 0)
}
func (wb *wb受払出荷) Exists(eb Eb受払出荷) Wb受払出荷 {
	wb.ebs = append(wb.ebs, eb)
	return wb
}
func (wb *wb受払出荷) build(argsCntStart ...int) (where Where) {
	where.w = ""
	where.prms = make([]interface{}, 0, len(wb.config))
	if len(argsCntStart) == 1 {
		where.argsCnt = argsCntStart[0]
	}
	for _, e := range wb.config {
		switch e.op {
		case OpIn:
			where.argsCnt++
			where.w += fmt.Sprintf(" AND (\"%s\"%s)", e.field, fmt.Sprintf(e.op.string(), fmt.Sprintf("$%d", where.argsCnt)))
			where.prms = append(where.prms, pq.Array(e.val))
			continue
		case OpIsNull:
			fallthrough
		case OpIsNotNull:
			where.w += fmt.Sprintf(" AND (\"%s\"%s)", e.field, e.op.string())
			continue
		default:
			where.argsCnt++
			where.w += fmt.Sprintf(" AND (\"%s\"%s)", e.field, fmt.Sprintf(e.op.string(), fmt.Sprintf("$%d", where.argsCnt)))
			where.prms = append(where.prms, e.val)
			continue
		}
	}
	for _, eb := range wb.ebs {
		w := eb.buildEb受払出荷(where.argsCnt)
		where.Append(w)
	}
	return
}

type nothingWb受払出荷 struct{}

func (wb *nothingWb受払出荷) And(field fld受払出荷, op whereBuilderOperater, val interface{}) Wb受払出荷 {
	return wb
}
func (wb *nothingWb受払出荷) Clear() {}
func (wb *nothingWb受払出荷) Exists(_ Eb受払出荷) Wb受払出荷 {
	return wb
}
func (wb *nothingWb受払出荷) build(argCntStart ...int) (where Where) {
	return Where{w: " AND 1<>1"}
}
func NewNothingWb受払出荷() Wb受払出荷 {
	return &nothingWb受払出荷{}
}
