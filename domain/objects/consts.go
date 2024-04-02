package objects

import (
	"errors"
)

var ErrArg error = errors.New("arg error")
var ErrNotFound error = errors.New("not found")
var ErrAlreadyExists error = errors.New("already exists")
var ErrOverMaxLength error = errors.New("over max length error")
