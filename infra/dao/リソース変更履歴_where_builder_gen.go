// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"fmt"
	pq "github.com/lib/pq"
)

type Wbリソース変更履歴 interface {
	And(field fldリソース変更履歴, op whereBuilderOperater, val interface{}) Wbリソース変更履歴
	Clear()
	build(argCntStart ...int) (where Where)
}
type wbリソース変更履歴 struct {
	config []whereBuilderExp
}

func NewWbリソース変更履歴() Wbリソース変更履歴 {
	return &wbリソース変更履歴{config: make([]whereBuilderExp, 0)}
}
func newWbリソース変更履歴WithPrimaryKeys(No Id) Wbリソース変更履歴 {
	wb := &wbリソース変更履歴{config: make([]whereBuilderExp, 0)}
	wb.And(Tblリソース変更履歴().FldNo(), OpEqu, No)

	return wb
}
func (wb *wbリソース変更履歴) And(field fldリソース変更履歴, op whereBuilderOperater, val interface{}) Wbリソース変更履歴 {
	wb.config = append(wb.config, whereBuilderExp{
		field: string(field),
		op:    op,
		val:   val,
	})
	return wb
}
func (wb *wbリソース変更履歴) Clear() {
	wb.config = make([]whereBuilderExp, 0)
}
func (wb *wbリソース変更履歴) build(argCntStart ...int) (where Where) {
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

type nothingWbリソース変更履歴 struct{}

func (wb *nothingWbリソース変更履歴) And(field fldリソース変更履歴, op whereBuilderOperater, val interface{}) Wbリソース変更履歴 {
	return wb
}
func (wb *nothingWbリソース変更履歴) Clear() {}
func (wb *nothingWbリソース変更履歴) build(argCntStart ...int) (where Where) {
	return Where{w: " AND 1<>1"}
}
func NewNothingWbリソース変更履歴() Wbリソース変更履歴 {
	return &nothingWbリソース変更履歴{}
}
