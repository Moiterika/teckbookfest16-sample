// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"fmt"
	pq "github.com/lib/pq"
)

type Wbログアップロード時 interface {
	And(field fldログアップロード時, op whereBuilderOperater, val interface{}) Wbログアップロード時
	Clear()
	Exists(...Ebログアップロード時)
	build(argCntStart ...int) (where Where)
}
type wbログアップロード時 struct {
	config []whereBuilderExp
	ebs    []Ebログアップロード時
}

func NewWbログアップロード時() Wbログアップロード時 {
	return &wbログアップロード時{
		config: make([]whereBuilderExp, 0),
		ebs:    make([]Ebログアップロード時, 0),
	}
}
func newWbログアップロード時WithPrimaryKeys(No Id) Wbログアップロード時 {
	wb := &wbログアップロード時{config: make([]whereBuilderExp, 0)}
	wb.And(Tblログアップロード時().FldNo(), OpEqu, No)

	return wb
}
func (wb *wbログアップロード時) And(field fldログアップロード時, op whereBuilderOperater, val interface{}) Wbログアップロード時 {
	wb.config = append(wb.config, whereBuilderExp{
		field: string(field),
		op:    op,
		val:   val,
	})
	return wb
}
func (wb *wbログアップロード時) Clear() {
	wb.config = make([]whereBuilderExp, 0)
}
func (wb *wbログアップロード時) Exists(ebs ...Ebログアップロード時) {
	wb.ebs = append(wb.ebs, ebs...)
}
func (wb *wbログアップロード時) build(argsCntStart ...int) (where Where) {
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
		w := eb.buildEbログアップロード時(where.argsCnt)
		where.Append(w)
	}
	return
}

type nothingWbログアップロード時 struct{}

func (wb *nothingWbログアップロード時) And(field fldログアップロード時, op whereBuilderOperater, val interface{}) Wbログアップロード時 {
	return wb
}
func (wb *nothingWbログアップロード時) Clear()                  {}
func (wb *nothingWbログアップロード時) Exists(_ ...Ebログアップロード時) {}
func (wb *nothingWbログアップロード時) build(argCntStart ...int) (where Where) {
	return Where{w: " AND 1<>1"}
}
func NewNothingWbログアップロード時() Wbログアップロード時 {
	return &nothingWbログアップロード時{}
}
