package dao

type Where struct {
	// Where句本体
	w string
	// パラメータ
	prms []interface{}
}

func (w Where) String() string {
	return w.w
}

func (w Where) Params() ([]interface{}, bool) {
	return w.prms, len(w.prms) > 0
}
