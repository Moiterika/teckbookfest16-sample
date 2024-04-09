package apl

import (
	"techbookfest16-sample/domain/types"

	"github.com/shopspring/decimal"
)

type Dto仕入品 struct {
	Getコード        types.Code品目      `json:"コード"`
	Get名称         string            `json:"名称"`
	Get基準単位コード    types.Code単位      `json:"基準単位コード"`
	Get生産用品目区分コード types.Code生産用品目区分 `json:"生産用品目区分コード"`
	Get標準単価       decimal.Decimal   `json:"標準単価"`
}

func (d Dto仕入品) 標準単価() types.Price {
	return types.NewPrice(d.Get標準単価, types.Jpy, d.Get基準単位コード)
}
