package objects

import "github.com/shopspring/decimal"

// アップロード履歴.処理状況
type Enum受払区分 int

const (
	Enum受払区分_仕入 Enum受払区分 = iota + 1
	Enum受払区分_仕入返品
	Enum受払区分_投入
	Enum受払区分_製造
	Enum受払区分_出荷
	Enum受払区分_出荷返品
	Enum受払区分_品目振替払出
	Enum受払区分_品目振替受入
	Enum受払区分_移動出
	Enum受払区分_移動入
	Enum受払区分_ロット振替出
	Enum受払区分_ロット振替入
	Enum受払区分_支給
	Enum受払区分_支給返品
	Enum受払区分_他勘定振替払出
	Enum受払区分_他勘定振替受入
)

func (e Enum受払区分) String() string {
	switch e {
	case Enum受払区分_仕入:
		return "仕入"
	case Enum受払区分_仕入返品:
		return "仕入返品"
	case Enum受払区分_投入:
		return "投入"
	case Enum受払区分_製造:
		return "製造"
	case Enum受払区分_出荷:
		return "出荷"
	case Enum受払区分_出荷返品:
		return "出荷返品"
	case Enum受払区分_品目振替払出:
		return "品目振替払出"
	case Enum受払区分_品目振替受入:
		return "品目振替受入"
	case Enum受払区分_移動出:
		return "移動出"
	case Enum受払区分_移動入:
		return "移動入"
	case Enum受払区分_ロット振替出:
		return "ロット振替出"
	case Enum受払区分_ロット振替入:
		return "ロット振替入"
	case Enum受払区分_支給:
		return "支給"
	case Enum受払区分_支給返品:
		return "支給返品"
	case Enum受払区分_他勘定振替払出:
		return "他勘定振替払出"
	case Enum受払区分_他勘定振替受入:
		return "他勘定振替受入"
	default:
		return "未定義"
	}
}

func (e Enum受払区分) Is入庫() bool {
	switch e {
	case Enum受払区分_仕入:
		fallthrough
	case Enum受払区分_仕入返品:
		fallthrough
	case Enum受払区分_製造:
		fallthrough
	case Enum受払区分_品目振替受入:
		fallthrough
	case Enum受払区分_移動入:
		fallthrough
	case Enum受払区分_ロット振替入:
		fallthrough
	case Enum受払区分_他勘定振替受入:
		return true

	case Enum受払区分_投入:
		fallthrough
	case Enum受払区分_出荷:
		fallthrough
	case Enum受払区分_出荷返品:
		fallthrough
	case Enum受払区分_品目振替払出:
		fallthrough
	case Enum受払区分_移動出:
		fallthrough
	case Enum受払区分_ロット振替出:
		fallthrough
	case Enum受払区分_支給:
		fallthrough
	case Enum受払区分_支給返品:
		fallthrough
	case Enum受払区分_他勘定振替払出:
		return false
	default:
		return false
	}
}

func (e Enum受払区分) Op符号() decimal.Decimal {
	switch e {
	case Enum受払区分_仕入:
		fallthrough
	case Enum受払区分_製造:
		fallthrough
	case Enum受払区分_出荷返品:
		fallthrough
	case Enum受払区分_品目振替受入:
		fallthrough
	case Enum受払区分_支給返品:
		fallthrough
	case Enum受払区分_他勘定振替受入:
		return decimal.NewFromInt32(1)

	case Enum受払区分_仕入返品:
		fallthrough
	case Enum受払区分_投入:
		fallthrough
	case Enum受払区分_出荷:
		fallthrough
	case Enum受払区分_品目振替払出:
		fallthrough
	case Enum受払区分_支給:
		fallthrough
	case Enum受払区分_他勘定振替払出:
		return decimal.NewFromInt32(-1)

	case Enum受払区分_移動出:
		fallthrough
	case Enum受払区分_移動入:
		fallthrough
	case Enum受払区分_ロット振替出:
		fallthrough
	case Enum受払区分_ロット振替入:
		return decimal.Zero
	default:
		return decimal.Zero
	}
}
