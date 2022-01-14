package errors

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidPassword    = errors.New("password is invalid")
	ErrIncorrectPassword  = errors.New("password is incorrect")
	ErrUserDoesNotExist   = errors.New("user does not exist")
	ErrAccountBlacklisted = errors.New("account has been blacklisted")
	ErrAccountBlocked     = errors.New("account has been blocked")
	ErrAccountClosed      = errors.New("account is closed")
	ErrInvalidPermission  = errors.New("invalid permission")
)
