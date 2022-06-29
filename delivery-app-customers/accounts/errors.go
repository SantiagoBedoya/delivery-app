package accounts

import "errors"

var (
	// ErrDuplicateEmail defines error when account email is already registered
	ErrDuplicateEmail = errors.New("an account exists with this email")
	// ErrAccountNotFound defines error when account is not found by email
	ErrAccountNotFound = errors.New("this account doesn't exists")
)
