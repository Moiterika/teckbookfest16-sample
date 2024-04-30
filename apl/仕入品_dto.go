package apl

import (
	"techbookfest16-sample/domain/types"
)

type Dto仕入品 struct {
	Getコード        types.Code品目      `json:"コード"`
	Get名称         string            `json:"名称"`
	Get基準単位コード    types.Code単位      `json:"基準単位コード"`
	Get生産用品目区分コード types.Code生産用品目区分 `json:"生産用品目区分コード"`
	Get標準単価       types.Price       `json:"標準単価"`
}
