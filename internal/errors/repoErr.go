package errors

import "errors"

var (
	ErrUserInsertion          = errors.New("users insertion error")
	ErrGettingByID            = errors.New("get users by ID error")
	ErrGettingByEmail         = errors.New("get users by email error")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")

	ErrSelection = errors.New("selection error")
	ErrNotAdmin  = errors.New("not admin")

	ErrTokenInsertion    = errors.New("tokens insertion error")
	ErrTokenDeletion     = errors.New("tokens deletion error")
	ErrUpdatingLogoutPin = errors.New("update logout_pin error")
	ErrGettingByToken    = errors.New("get refresh tokens error")
)
