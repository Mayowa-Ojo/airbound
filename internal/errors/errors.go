package errors

import "errors"

var (
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrInvalidPassword        = errors.New("password is invalid")
	ErrIncorrectPassword      = errors.New("password is incorrect")
	ErrUserDoesNotExist       = errors.New("user does not exist")
	ErrAccountBlacklisted     = errors.New("account has been blacklisted")
	ErrAccountBlocked         = errors.New("account has been blocked")
	ErrAccountClosed          = errors.New("account is closed")
	ErrInvalidPermission      = errors.New("invalid permission")
	ErrAccountAlreadyVerified = errors.New("account is already verified")
	ErrTwoFaSetupCompleted    = errors.New("two-factor authentication setup is already completed")
	ErrTwoFaSetupIncomplete   = errors.New("two-factor authentication setup is not completed")
	ErrInvalidTwoFaCode       = errors.New("two-factor code is invalid/missing")
	ErrInvalidInternalEmail   = errors.New("email is not a valid internal domain")
)
