package system

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

var (
	ErrNotFound = errors.New("tag not found")
)

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func InsertTx(ctx context.Context, tx *sqlx.Tx, tag *System) (*System, error) {
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
                       null,
                       null
                   ) returning id,
                               name,
                               created_at,
                               updated_at,
                               deleted_at;
`
	var g System
	rows, err := tx.NamedQuery(query, tag)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) FindByPattern(ctx context.Context, pattern string) ([]System, error) {
	pattern = "%" + pattern + "%"
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from tags
	   where deleted_at is null
		 and name like $1;
`
	groups := make([]System, 0)
	err := r.conn.SelectContext(ctx, &groups, query, pattern)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func GetByNameDb(ctx context.Context, conn *sqlx.DB, name string) (*System, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from tags
	   where deleted_at is null
		 and name = $1;
`
	var tag System
	err := conn.GetContext(ctx, &tag, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &tag, nil
}

func (r *repository) List(ctx context.Context) ([]System, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from systems
	   where deleted_at is null;
`
	groups := make([]System, 0)
	err := r.conn.SelectContext(ctx, &groups, query)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

