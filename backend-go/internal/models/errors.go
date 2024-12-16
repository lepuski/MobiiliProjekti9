package models

import (
	"errors"
)

var (
	ErrNoRecord = errors.New("no matching record found")

	ErrInvalidCredentials = errors.New("models: invalid credentials")

	ErrDuplicateUsername = errors.New("models: duplicate name")

	ErrUserNotFound = errors.New("user not found")
)
