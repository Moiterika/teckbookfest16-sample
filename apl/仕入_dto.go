package apl

import (
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"time"

	"github.com/shopspring/decimal"
)

type Dto仕入 struct {
	Get登録日時   time.Time          `json:"登録日時"`
	Get計上月    time.Time          `json:"計上月"`
	Get受払区分   objects.Enum受払区分   `json:"受払区分"`
	Get赤伝フラグ  bool               `json:"赤伝フラグ"`
	Get品目コード  types.Code品目       `json:"品目コード"`
	Get基準数量   decimal.Decimal    `json:"基準数量"`
	Get基準数量単位 types.Code単位       `json:"基準数量単位"`
	Get仕入数量   decimal.Decimal    `json:"仕入数量"`
	Get仕入数量単位 types.Code単位       `json:"仕入数量単位"`
	Get仕入金額   decimal.Decimal    `json:"仕入金額"`
	Get仕入金額通貨 types.CurrencyUnit `json:"仕入金額通貨"`
	Get仕入単価   decimal.Decimal    `json:"仕入単価"`
}

func (d Dto仕入) 基準数量() types.Inventory {
	return types.NewInventory(d.Get基準数量, d.Get基準数量単位)
}
func (d Dto仕入) 仕入数量() (types.Quantity, error) {
	return types.NewQuantity(d.Get仕入数量, d.Get仕入数量単位)
}
func (d Dto仕入) 仕入金額() types.Amount {
	return types.NewAmount(d.Get仕入金額, d.Get仕入金額通貨)
}
func (d Dto仕入) 仕入単価() types.Price {
	return types.NewPrice(d.Get仕入単価, d.Get仕入金額通貨, d.Get仕入数量単位)
}
