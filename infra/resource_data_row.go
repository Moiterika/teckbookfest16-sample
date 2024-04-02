package infra

import "techbookfest16-sample/infra/dao"

type ResourceDataRow interface {
	TableName() string
	RowState() dao.DataRowState
	ToJson() ([]byte, error)
}
