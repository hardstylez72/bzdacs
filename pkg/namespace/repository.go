package namespace

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}
func (r *repository) Update(ctx context.Context, namespace *Namespace) (*Namespace, error) {
	return Update(ctx, r.conn, namespace)
}

func Update(ctx context.Context, conn storage.SqlDriver, namespace *Namespace) (*Namespace, error) {
	query := `
	update namespaces
	   set name = :name,
		   updated_at = now()
	 where id = :id returning  id,
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
			return nil, storage.PgError(err)
		}
	}

	if g.Id == 0 {
		return nil, storage.EntityNotFound
	}
	return &g, nil
}

func (r *repository) Delete(ctx context.Context, namespaceId int) error {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	txx := storage.WrapSqlxTx(tx)

	err = Delete(ctx, txx, namespaceId)
	if err != nil {
		return err
	}
	return nil
}

func Delete(ctx context.Context, conn storage.SqlDriver, namespaceId int) error {

	query := `
			update namespaces
				set deleted_at = now()
				where id = $1;
`
	_, err := conn.ExecContext(ctx, query, namespaceId)
	if err != nil {
		return storage.PgError(err)
	}

	return nil
}
