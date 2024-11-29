package models

import "errors"

var ErrNotFound = errors.New("item not found")

var ErrConflict = errors.New("item already exists")
