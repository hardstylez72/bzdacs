package util

import "errors"

var (
	ErrEntityAlreadyExists = errors.New("entity already exist")
	ErrEntityNotFound      = errors.New("entity not found")
)
