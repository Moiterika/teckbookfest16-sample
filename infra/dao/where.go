package dao

type Where struct {
	// Where句本体
	w string
	// パラメータ
	prms []interface{}
	// パラメータ現在値
	argsCnt int
}

func (w Where) String() string {
	return w.w
}

func (w Where) Params() ([]interface{}, bool) {
	return w.prms, len(w.prms) > 0
}

func (w *Where) Append(ws ...Where) {
	for _, e := range ws {
		w.w += e.w
		w.prms = append(w.prms, e.prms...)
		w.argsCnt += e.argsCnt
	}
}

// func (w *Where) Concatf(e Where) {
// 	w.w = fmt.Sprintf(w.w, e.w)
// 	w.prms = append(w.prms, e.prms...)
// 	w.argsCnt += e.argsCnt
// }
