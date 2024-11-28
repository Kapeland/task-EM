package models

import "errors"

var ErrNotFound = errors.New("item not found")

var ErrInvalidInput = errors.New("invalid input")

var ErrConflict = errors.New("item already exists")
