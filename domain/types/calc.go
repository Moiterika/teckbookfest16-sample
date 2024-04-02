package types

import (
	"github.com/shopspring/decimal"
	"golang.org/x/xerrors"
)

// 単価の丸め位置初期値（小数点以下第6位までを保持。※7位で四捨五入）
const DefaultPriceRound int32 = 6

func SumAmount(x ...Amount) (ret Amount, err error) {
	if len(x) == 0 {
		err = xerrors.Errorf("計算不能")
		return
	}

	// 初期値
	ret.val = decimal.Zero
	ret.unit = x[0].unit
	curMap := make(map[CurrencyUnit]struct{})
	for i := range x {
		if _, ok := curMap[x[i].unit]; !ok {
			err = xerrors.Errorf("単位違うの居る……")
			return
		}
		ret.val = ret.val.Add(x[i].val)
	}
	return
}

func SumQuantity(x ...Quantity) (ret Quantity, err error) {
	if len(x) == 0 {
		err = xerrors.Errorf("計算不能")
		return
	}

	// 初期値
	ret.val = decimal.Zero
	ret.unit = x[0].unit
	unitMap := make(map[Code単位]struct{})
	for i := range x {
		if _, ok := unitMap[x[i].unit]; !ok {
			err = xerrors.Errorf("単位違うの居る……")
			return
		}
		ret.val = ret.val.Add(x[i].val)
	}
	return
}

func SumInventory(x ...Inventory) (ret Inventory, err error) {
	if len(x) == 0 {
		err = xerrors.Errorf("計算不能")
		return
	}

	// 初期値
	ret.val = decimal.Zero
	ret.unit = x[0].unit
	unitMap := make(map[Code単位]struct{})
	for i := range x {
		if _, ok := unitMap[x[i].unit]; !ok {
			err = xerrors.Errorf("単位違うの居る……")
			return
		}
		ret.val = ret.val.Add(x[i].val)
	}
	return
}

func Calc単価(a Amount, q Quantity) (p Price, err error) {
	if q.val.Equal(decimal.Zero) {
		err = xerrors.Errorf("0除算エラー: 金額=%d、数量=%d", a.val, q.val)
	}
	p.amt = a.val.Div(q.val).Round(6)
	p.cur = a.unit
	p.perUnit = q.unit
	return
}

func Calc金額(q Quantity, p Price) (a Amount, err error) {
	if q.unit != p.perUnit {
		err = xerrors.Errorf("計算不能")
		return
	}

	a = NewAmount(q.val.Mul(p.amt), p.cur)
	return
}

type ProrationBasis[T Code単位 | CurrencyUnit] interface {
	Val() decimal.Decimal
	Unit() T
}

// Prorate は金額aをbsで按分する。 returns:
//
//	 as: 按分後金額のリスト
//	err: エラー
func Prorate[T Code単位 | CurrencyUnit](a Amount, bs []ProrationBasis[T]) (as []Amount, err error) {
	if len(bs) == 0 {
		err = xerrors.Errorf("計算不能")
		return
	}

	basisTotal := decimal.Zero
	maxIndex := 0
	tempMax := bs[0].Val()
	unitMap := map[T]struct{}{bs[0].Unit(): {}}
	for i, b := range bs {
		// 単位チェック
		if _, ok := unitMap[b.Unit()]; !ok {
			err = xerrors.Errorf("単位違うの居る……")
			return
		}
		basisTotal = basisTotal.Add(b.Val())
		// tmpMax < y ⇒ yをtmpMaxにする
		if tempMax.Cmp(b.Val()) == -1 {
			maxIndex = i
			tempMax = b.Val()
		}
	}

	if basisTotal.Equal(decimal.Zero) {
		err = xerrors.Errorf("按分できない……")
		return
	}

	as = make([]Amount, len(bs))
	sumAs := decimal.Zero
	for i, b := range bs {
		as[i] = Amount{
			val:  a.val.Mul(b.Val()).Div(basisTotal).Round(0),
			unit: a.unit,
		}
		sumAs = sumAs.Add(as[i].val)
	}
	// 端数調整
	diff := a.val.Add(sumAs.Neg())
	if !diff.Equal(decimal.Zero) {
		as[maxIndex].val = as[maxIndex].val.Add(diff)
	}
	return
}

// Calc按分 は按分元金額を按分基準一覧で按分する。
func Calc按分[T Code単位 | CurrencyUnit](按分元金額 Amount, 按分基準一覧 []ProrationBasis[T]) (按分結果 []Amount, err error) {
	if len(按分基準一覧) == 0 {
		err = xerrors.Errorf("計算不能")
		return
	}

	按分基準合計 := decimal.Zero
	最大値index := 0
	最大値 := 按分基準一覧[0].Val()
	unitMap := map[T]struct{}{按分基準一覧[0].Unit(): {}}
	for i, b := range 按分基準一覧 {
		// 単位チェック
		if _, ok := unitMap[b.Unit()]; !ok {
			err = xerrors.Errorf("単位違うの居る……")
			return
		}
		按分基準合計 = 按分基準合計.Add(b.Val())
		// tmpMax < y ⇒ yをtmpMaxにする
		if 最大値.Cmp(b.Val()) == -1 {
			最大値index = i
			最大値 = b.Val()
		}
	}

	if 按分基準合計.Equal(decimal.Zero) {
		err = xerrors.Errorf("按分できない……")
		return
	}

	按分結果 = make([]Amount, len(按分基準一覧))
	按分結果合計 := decimal.Zero
	for i, 按分基準 := range 按分基準一覧 {
		按分結果[i] = Amount{
			val:  按分元金額.val.Mul(按分基準.Val()).Div(按分基準合計).Round(0),
			unit: 按分元金額.unit,
		}
		按分結果合計 = 按分結果合計.Add(按分結果[i].val)
	}
	// 端数調整
	端数 := 按分元金額.val.Add(按分結果合計.Neg())
	if !端数.Equal(decimal.Zero) {
		按分結果[最大値index].val = 按分結果[最大値index].val.Add(端数)
	}
	return
}

// MulQRate returns q * rate
func MulQRate(q Quantity, rate Rate) Quantity {
	return Quantity{
		val:  q.val.Mul(decimal.Decimal(rate)),
		unit: q.unit,
	}
}

// MulQPrice returns q * price
func MulQPrice(q Quantity, price Price) (amt Amount, err error) {
	if q.unit != price.perUnit {
		err = xerrors.Errorf("単位違うし")
		return
	}

	amt.val = q.val.Mul(price.amt)
	amt.unit = price.cur
	return
}

// MulARate returns a * rate
func MulARate(a Amount, rate Rate, round ...int32) (amt Amount) {
	var n int32 = 0
	if len(round) == 1 {
		n = round[0]
	}
	amt.val = a.val.Mul(decimal.Decimal(rate)).Round(n)
	amt.unit = a.unit
	return
}

// CalcPrice returns a / q rounded with DefaultPriceRound.
func CalcPrice(a Amount, q Quantity) (p Price) {
	if q.IsZero() {
		p.amt = decimal.Zero
		p.cur = a.unit
		p.perUnit = q.unit
		return
	}

	p.amt = a.val.Div(q.val).Round(DefaultPriceRound)
	p.cur = a.unit
	p.perUnit = q.unit
	return
}
