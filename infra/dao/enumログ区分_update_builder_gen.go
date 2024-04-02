// Code generated by xlsx2go.go; DO NOT EDIT.

package dao

import (
	"fmt"
	"strings"
)

type ubEnumログ区分 struct {
	config          []updateBuilderKvPair
	keys            map[string]int
	isBuildComplete bool
}

func NewUbEnumログ区分() *ubEnumログ区分 {
	return &ubEnumログ区分{
		config:          make([]updateBuilderKvPair, 0),
		isBuildComplete: false,
		keys:            make(map[string]int),
	}
}
func (ub *ubEnumログ区分) Set(field fldEnumログ区分, val interface{}) *ubEnumログ区分 {
	k := string(field)
	index, exists := ub.keys[k]
	if exists {
		// 同じキーに対してSetされた場合、値は上書きされる
		ub.config[index].val = val
		return ub
	}
	ub.config = append(ub.config, updateBuilderKvPair{
		key: k,
		val: val,
	})
	ub.keys[k] = len(ub.config) - 1
	return ub
}
func (ub *ubEnumログ区分) Clear() {
	ub.config = make([]updateBuilderKvPair, 0)
	ub.keys = make(map[string]int)
	ub.isBuildComplete = false
}
func (ub *ubEnumログ区分) Count() int {
	return len(ub.config)
}
func (ub *ubEnumログ区分) build(wb WbEnumログ区分) (s string, w string, execArgs []interface{}) {
	if ub.Count() == 0 {
		// 更新項目なし
		return
	}
	where := wb.build(ub.Count() + 1)
	whereParams, exists := where.Params()
	execArgs = make([]interface{}, ub.Count(), ub.Count()+len(whereParams))
	tmp := make([]string, ub.Count())
	for i, v := range ub.config {
		tmp[i] = fmt.Sprintf(" \"%s\" = $%d", v.key, i+1)
		execArgs[i] = v.val
	}
	if exists {
		execArgs = append(execArgs, whereParams...)
	}
	s = strings.Join(tmp, ",")
	ub.isBuildComplete = true
	return
}
func (ub *ubEnumログ区分) copyMap() map[string]interface{} {
	ret := make(map[string]interface{})
	for _, v := range ub.config {
		ret[v.key] = v.val
	}
	return ret
}
