// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import "fmt"

type ebログアップロード時joinログ struct {
	wb Wbログ
}

func NewEbログアップロード時joinログ() *ebログアップロード時joinログ {
	return &ebログアップロード時joinログ{wb: NewWbログ()}
}
func (eb *ebログアップロード時joinログ) And(field fldログ, op whereBuilderOperater, val interface{}) Ebログアップロード時 {
	eb.wb.And(field, op, val)
	return eb
}
func (eb *ebログアップロード時joinログ) buildEbログアップロード時(argCntStart ...int) (where Where) {
	where = eb.wb.build(argCntStart...)
	where.w = fmt.Sprintf(" AND EXISTS (SELECT * FROM \"ログ\" WHERE \"ログ\".\"No\" = \"ログ_アップロード時\".\"No\" %s)", where.w)
	return
}
