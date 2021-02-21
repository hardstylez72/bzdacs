package storage

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
)

const (
	UniqExceptionCode = "23505"
)

var (
	EntityNotFound     = errors.New("entity not found")
	EntityAlreadyExist = errors.New("entity with same attribute already exist")
)

func PgError(err error) error {

	if sql.ErrNoRows == err {
		return EntityNotFound
	}

	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		if pgErr.Code == UniqExceptionCode {
			return EntityAlreadyExist
		}
	}

	return nil
}
