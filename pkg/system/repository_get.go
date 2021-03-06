package system

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"strconv"
)

const (
	noName = ""
	noId   = 0
)

const (
	GetById   GetKind = "GetById"
	GetByName GetKind = "GetByName"
)

func (r *repository) Get(ctx context.Context, kind GetKind, arg ...interface{}) (*System, error) {

	if len(arg) != 1 {
		return nil, errors.New("invalid amount of input params: " + strconv.Itoa(len(arg)) + " expected 1 param")
	}
	param := arg[0]

	switch kind {
	case GetById:
		id, ok := param.(int)
		if !ok {
			return nil, errors.New("invalid id param, must be type int, but have: ")
		}
		return Get(ctx, r.conn, id, noName)
	case GetByName:
		name, ok := param.(string)
		if !ok {
			return nil, errors.New("invalid id param, must be type string, but have: ")
		}
		return Get(ctx, r.conn, noId, name)
	}
	return nil, errors.New("invalid input params")
}

func Get(ctx context.Context, driver storage.SqlDriver, id int, name string) (*System, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Select(`
 			   id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
`).From("systems")

	if id != 0 {
		builder = builder.Where(sq.Eq{"id": id})
	}

	if name != "" {
		builder = builder.Where(sq.Eq{"name": name})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	query += ";"

	var system System
	err = driver.GetContext(ctx, &system, query, args...)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &system, nil
}
