// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import "fmt"

type eb受払join受払出荷 struct {
	wb Wb受払出荷
}

func NewEb受払join受払出荷() *eb受払join受払出荷 {
	return &eb受払join受払出荷{wb: NewWb受払出荷()}
}
func (eb *eb受払join受払出荷) And(field fld受払出荷, op whereBuilderOperater, val interface{}) Eb受払 {
	eb.wb.And(field, op, val)
	return eb
}
func (eb *eb受払join受払出荷) buildEb受払(argCntStart ...int) (where Where) {
	where = eb.wb.build(argCntStart...)
	where.w = fmt.Sprintf(" AND EXISTS (SELECT * FROM \"受払_出荷\" WHERE \"受払_出荷\".\"No\" = \"受払\".\"No\" %s)", where.w)
	return
}
