package dao

type DataRowState int8

const (
	Detached DataRowState = 1 << iota
	UnChanged
	Added
	Deleted
	Modified
)

func (d DataRowState) String() string {
	switch d {
	case Added:
		return "A"
	case Modified:
		return "M"
	case Deleted:
		return "D"
	default:
		return "-"
	}
}
