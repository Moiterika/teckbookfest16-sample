// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import "fmt"

type eb品目join単位 struct {
	wb Wb単位
}

func NewEb品目join単位() Eb品目 {
	return &eb品目join単位{wb: NewWb単位()}
}
func (eb *eb品目join単位) And(field fld単位, op whereBuilderOperater, val interface{}) Eb品目 {
	eb.wb.And(field, op, val)
	return eb
}
func (eb *eb品目join単位) buildEb品目(argCntStart ...int) (where Where) {
	where = eb.wb.build(argCntStart...)
	where.w = fmt.Sprintf(" AND EXISTS (SELECT * FROM \"単位\" WHERE \"単位\".\"ID\" = \"品目\".\"基準単位ID\" %s)", where.w)
	return
}
