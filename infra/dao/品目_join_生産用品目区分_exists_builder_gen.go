// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import "fmt"

type eb品目join生産用品目区分 struct {
	wb Wb生産用品目区分
}

func NewEb品目join生産用品目区分() *eb品目join生産用品目区分 {
	return &eb品目join生産用品目区分{wb: NewWb生産用品目区分()}
}
func (eb *eb品目join生産用品目区分) And(field fld生産用品目区分, op whereBuilderOperater, val interface{}) Eb品目 {
	eb.wb.And(field, op, val)
	return eb
}
func (eb *eb品目join生産用品目区分) buildEb品目(argCntStart ...int) (where Where) {
	where = eb.wb.build(argCntStart...)
	where.w = fmt.Sprintf(" AND EXISTS (SELECT * FROM \"生産用品目区分\" WHERE \"生産用品目区分\".\"ID\" = \"品目\".\"生産用品目区分ID\" %s)", where.w)
	return
}
