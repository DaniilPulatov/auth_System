package errors

import "errors"

var (
	ErrNewEnv = errors.New("error loading env")
)
