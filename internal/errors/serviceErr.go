package errors

import "errors"

var (
	ErrPasswordHashing = errors.New("password hashing failed")
)
