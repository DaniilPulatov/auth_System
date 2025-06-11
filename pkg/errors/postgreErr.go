package errors

import "errors"

var (
	ErrParseConfig    = errors.New("error parsing config")
	ErrNewWithConfig  = errors.New("error new config")
	ErrPingConnection = errors.New("error ping connection")
)
