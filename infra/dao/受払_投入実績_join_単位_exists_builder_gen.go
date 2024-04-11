// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import "fmt"

type eb受払投入実績join単位 struct {
	wb Wb単位
}

func NewEb受払投入実績join単位() Eb受払投入実績 {
	return &eb受払投入実績join単位{wb: NewWb単位()}
}
func (eb *eb受払投入実績join単位) And(field fld単位, op whereBuilderOperater, val interface{}) Eb受払投入実績 {
	eb.wb.And(field, op, val)
	return eb
}
func (eb *eb受払投入実績join単位) buildEb受払投入実績(argCntStart ...int) (where Where) {
	where = eb.wb.build(argCntStart...)
	where.w = fmt.Sprintf(" AND EXISTS (SELECT * FROM \"単位\" WHERE \"単位\".\"ID\" = \"受払_投入実績\".\"投入単位ID\" %s)", where.w)
	return
}
