package error

import (
	"errors"
)

var (
	BOOL_ERR  = errors.New("error converting to boolean")
	FLOAT_ERR = errors.New("error converting to float")
	INT_ERR   = errors.New("error converting to integer")
)
