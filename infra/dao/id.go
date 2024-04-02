package dao

import (
	"strconv"
)

type Id int64

func (id Id) String() string {
	return strconv.FormatInt(int64(id), 10)
}

func (id Id) Unit() string {
	return id.String()
}
