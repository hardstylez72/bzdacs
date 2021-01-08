package tag

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

func (r *repository) Update(ctx context.Context, group *Tag) (*Tag, error) {
	query := `
	update ad.tags 
	   set name = :name,
		   updated_at = now()
	 where id = :id returning  id,
                               name,
                               created_at,
                               updated_at,
                               deleted_at;
`

	rows, err := r.conn.NamedQueryContext(ctx, query, group)
	if err != nil {
		return nil, err
	}

	var g Tag
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) Insert(ctx context.Context, group *Tag) (*Tag, error) {
	query := `
insert into ad.tags (
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

	rows, err := r.conn.NamedQueryContext(ctx, query, group)
	if err != nil {
		return nil, err
	}

	var g Tag
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func InsertTx(ctx context.Context, tx *sqlx.Tx, tag *Tag) (*Tag, error) {
	query := `
insert into ad.tags (
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
	var g Tag
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

func (r *repository) GetById(ctx context.Context, id int) (*Tag, error) {
	return GetByIdDb(ctx, r.conn, id)
}

func GetByIdDb(ctx context.Context, conn *sqlx.DB, id int) (*Tag, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from ad.tags
	   where id = $1
`
	var group Tag
	err := conn.GetContext(ctx, &group, query, id)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *repository) FindByPattern(ctx context.Context, pattern string) ([]Tag, error) {
	pattern = "%" + pattern + "%"
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from ad.tags
	   where deleted_at is null
		 and name like $1;
`
	groups := make([]Tag, 0)
	err := r.conn.SelectContext(ctx, &groups, query, pattern)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func GetByNameDb(ctx context.Context, conn *sqlx.DB, name string) (*Tag, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from ad.tags
	   where deleted_at is null
		 and name = $1;
`
	var tag Tag
	err := conn.GetContext(ctx, &tag, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &tag, nil
}

func (r *repository) List(ctx context.Context) ([]Tag, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from ad.tags
	   where deleted_at is null;
`
	groups := make([]Tag, 0)
	err := r.conn.SelectContext(ctx, &groups, query)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `
		update ad.tags 
			set deleted_at = now()
		where id = $1;
`
	_, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
