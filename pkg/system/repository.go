package system

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/namespace"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) Update(ctx context.Context, namespace *System) (*System, error) {
	return Update(ctx, r.conn, namespace)
}

func Update(ctx context.Context, conn *sqlx.DB, namespace *System) (*System, error) {
	query := `
	update systems 
	   set name = :name,
		   updated_at = now()
	 where id = :id returning  id,
                               name,
                               created_at,
                               updated_at,
                               deleted_at;
`

	rows, err := conn.NamedQueryContext(ctx, query, namespace)
	if err != nil {
		return nil, storage.PgError(err)
	}

	var g System
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, storage.PgError(err)
		}
	}

	return &g, nil
}

func (r *repository) Insert(ctx context.Context, system *System) (*System, error) {
	return Insert(ctx, r.conn, system)
}

func Insert(ctx context.Context, driver storage.SqlDriver, system *System) (*System, error) {
	query := `
insert into systems (
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

	rows, err := driver.NamedQueryContext(ctx, query, system)
	if err != nil {
		return nil, storage.PgError(err)
	}

	var g System
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, storage.PgError(err)
		}
	}

	return &g, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {

	namespaces, err := namespace.List(ctx, r.conn, id)
	if err != nil {
		return err
	}

	for _, ns := range namespaces {
		err = namespace.Delete(ctx, r.conn, ns.Id)
		if err != nil {
			return err
		}
	}
	return Delete(ctx, r.conn, id)
}

func Delete(ctx context.Context, conn *sqlx.DB, id int) error {
	query := `
		update systems 
			set deleted_at = now()
		where id = $1;
`
	_, err := conn.ExecContext(ctx, query, id)
	if err != nil {
		return storage.PgError(err)
	}

	return nil
}
