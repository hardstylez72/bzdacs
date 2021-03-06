package namespace

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

func (r *repository) Get(ctx context.Context, kind GetKind, systemId int, arg ...interface{}) (*Namespace, error) {
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
		return Get(ctx, r.conn, systemId, id, noName)
	case GetByName:
		name, ok := param.(string)
		if !ok {
			return nil, errors.New("invalid id param, must be type string, but have: ")
		}
		return Get(ctx, r.conn, systemId, noId, name)
	}
	return nil, errors.New("invalid input params")
}

func Get(ctx context.Context, conn storage.SqlDriver, systemId, namespaceId int, name string) (*Namespace, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Select(`
 			   n.id,
			   n.name,
			   n.created_at,
			   n.updated_at,
			   n.deleted_at,
			   n.system_id
			`).From("namespaces n")

	if namespaceId != 0 {
		builder = builder.Where(sq.Eq{"n.id": namespaceId})
	} else {
		if name != "" {
			builder = builder.Where(sq.Eq{"n.name": name})
		}
		builder = builder.Where(sq.Eq{"n.system_id": systemId})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var namespace Namespace
	err = conn.GetContext(ctx, &namespace, query, args...)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &namespace, nil
}
