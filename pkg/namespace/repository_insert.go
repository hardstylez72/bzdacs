package namespace

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/systemnamespace"
	"github.com/jackc/pgconn"
)

func (r *repository) Insert(ctx context.Context, namespace *NamespaceExt) (*Namespace, error) {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	txx := storage.WrapSqlxTx(tx)

	n, err := InsertConn(ctx, txx, &namespace.Namespace)
	if err != nil {
		return nil, err
	}

	err = systemnamespace.InsertConn(ctx, txx, systemnamespace.Pair{
		SystemId:    namespace.SystemId,
		NamespaceId: n.Id,
	})

	if err != nil {
		return nil, err
	}

	return n, nil
}

func InsertConn(ctx context.Context, conn storage.SqlDriver, namespace *Namespace) (*Namespace, error) {
	query := `
insert into namespaces (
                       name,
                       created_at,
                       updated_at,
                       deleted_at
                       )
                   values (
                       :name,
                       now(),
                       now(),
                       null
                   ) returning id,
                               name,
                               created_at,
                               updated_at,
                               deleted_at;
`

	rows, err := conn.NamedQueryContext(ctx, query, namespace)
	if err != nil {
		const UniqExceptionCode = "23505"
		pgErr, ok := err.(*pgconn.PgError)
		if ok {
			if pgErr.Code == UniqExceptionCode {
				return nil, errors.New("entity with same attribute already exist")
			}
		}
		return nil, err
	}

	var g Namespace
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}
