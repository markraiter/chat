package util

import "errors"

var (
	ErrEmailExist       = errors.New("this email is already exist")
	ErrUsernameExist    = errors.New("this username is already exist")
	ErrWrongCredentials = errors.New("invalid login or password")
)
