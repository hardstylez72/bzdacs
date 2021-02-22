package namespace

import (
	"context"
	sq "github.com/Masterminds/squirrel"
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
	return UpdateLL(ctx, r.conn, namespace)
}

func UpdateLL(ctx context.Context, conn storage.SqlDriver, namespace *Namespace) (*Namespace, error) {
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

func (r *repository) Get(ctx context.Context, systemId, namespaceId int, name string) (*Namespace, error) {
	return GetLL(ctx, r.conn, systemId, namespaceId, name)
}

func GetLL(ctx context.Context, conn storage.SqlDriver, systemId, namespaceId int, name string) (*Namespace, error) {
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

	err = DeleteLL(ctx, txx, namespaceId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLL(ctx context.Context, conn storage.SqlDriver, namespaceId int) error {

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
