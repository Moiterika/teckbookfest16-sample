package dao

type whereBuilderExp struct {
	field string
	op    whereBuilderOperater
	val   interface{}
}
