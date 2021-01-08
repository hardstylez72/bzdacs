package util

import "errors"

var ErrorMap map[error]string

var (
	ErrEntityAlreadyExists = errors.New("entity already exist")
	ErrEntityNotFound      = errors.New("entity not found")
)

func init() {
	ErrorMap = make(map[error]string)
	ErrorMap[ErrEntityAlreadyExists] = "entity already exist"
	ErrorMap[ErrEntityNotFound] = "entity not found"
}

func GetErrorCode(err error) string {
	code, ok := ErrorMap[err]
	if ok {
		return code
	}

	return "unknown error"
}
