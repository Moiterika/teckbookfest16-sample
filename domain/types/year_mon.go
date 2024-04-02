package types

import "time"

// YearMon は月度を表す。たとえば、計上月などに使う。
type YearMon struct {
	t time.Time
}

func NewYearMon(t time.Time) YearMon {
	return YearMon{
		t: t,
	}
}

func (ym *YearMon) Equals(other YearMon) bool {
	return ym.t.Year() == other.t.Year() && ym.t.Month() == other.t.Month()
}
