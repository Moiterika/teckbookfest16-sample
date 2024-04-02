package dao

type whereBuilderOperater int32

const (
	// ＝
	OpEqu whereBuilderOperater = iota
	// ≠
	OpNeq
	// ＞
	OpGtr
	// op_les                 // ＜
	// op_leq                 // ≦
	// op_geq                 // ≧
	// IN句
	OpIn
	// op_nin                 // NOT IN ()
	// op_lik                 // LIKE
	// op_nlk                 // NOT LIKE
	OpIsNotNull
	OpIsNull
)

func (o whereBuilderOperater) string() (s string) {
	switch o {
	case OpEqu:
		s = " = %s"
	case OpNeq:
		s = " <> %s"
	case OpGtr:
		s = " > %s"
	case OpIn:
		s = " = ANY(%s)"
	case OpIsNotNull:
		s = " IS NOT NULL"
	case OpIsNull:
		s = " IS NULL"
	}
	return
}
