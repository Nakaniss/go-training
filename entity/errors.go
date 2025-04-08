package entity

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidEmail = errors.New("invalid email")
)
