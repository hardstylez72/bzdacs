package namespace

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

func (r *repository) Insert(ctx context.Context, namespace *Namespace) (*Namespace, error) {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	txx := storage.WrapSqlxTx(tx)

	n, err := Insert(ctx, txx, namespace)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func Insert(ctx context.Context, conn storage.SqlDriver, namespace *Namespace) (*Namespace, error) {
	query := `
insert into namespaces (
                       name,
                       created_at,
                       updated_at,
                       deleted_at,
					   system_id
                       )
                   values (
                       :name,
                       now(),
                       now(),
                       null,
					   :system_id
                   ) returning id,
                               name,
                               created_at,
                               updated_at,
                               deleted_at,
							   system_id
`

	rows, err := conn.NamedQueryContext(ctx, query, namespace)
	if err != nil {
		return nil, storage.PgError(err)
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
