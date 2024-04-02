package types

// CurrencyUnit 型は通貨単位を表す。
type CurrencyUnit int16

const (
	Jpy CurrencyUnit = iota + 1
)

func (c CurrencyUnit) String() string {
	switch c {
	case Jpy:
		return "円"
	default:
		return "未定義"
	}
}
