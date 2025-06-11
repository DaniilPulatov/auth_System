package errors

import "errors"

var (
	ErrNewMigration = errors.New("new migration error")
	ErrUpMigrations = errors.New("up migrations error")
)
