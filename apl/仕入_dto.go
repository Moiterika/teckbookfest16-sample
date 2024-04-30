package apl

import (
	"techbookfest16-sample/domain/objects"
	"techbookfest16-sample/domain/types"
	"time"
)

type Dto仕入 struct {
	Get登録日時  time.Time        `json:"登録日時"`
	Get計上月   time.Time        `json:"計上月"`
	Get受払区分  objects.Enum受払区分 `json:"受払区分"`
	Get赤伝フラグ bool             `json:"赤伝フラグ"`
	Get品目コード types.Code品目     `json:"品目コード"`
	Get基準数量  types.Inventory  `json:"基準数量"`
	Get仕入数量  types.Quantity   `json:"仕入数量"`
	Get仕入金額  types.Amount     `json:"仕入金額"`
	Get仕入単価  types.Price      `json:"仕入単価"`
}
