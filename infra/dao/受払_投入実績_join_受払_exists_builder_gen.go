// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import "fmt"

type eb受払投入実績join受払 struct {
	wb Wb受払
}

func NewEb受払投入実績join受払() *eb受払投入実績join受払 {
	return &eb受払投入実績join受払{wb: NewWb受払()}
}
func (eb *eb受払投入実績join受払) And(field fld受払, op whereBuilderOperater, val interface{}) Eb受払投入実績 {
	eb.wb.And(field, op, val)
	return eb
}
func (eb *eb受払投入実績join受払) buildEb受払投入実績(argCntStart ...int) (where Where) {
	where = eb.wb.build(argCntStart...)
	where.w = fmt.Sprintf(" AND EXISTS (SELECT * FROM \"受払\" WHERE \"受払\".\"No\" = \"受払_投入実績\".\"No\" %s)", where.w)
	return
}
