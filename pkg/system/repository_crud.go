package system

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func (r *repository) Update(ctx context.Context, namespace *System) (*System, error) {
	return UpdateConn(ctx, r.conn, namespace)
}

func UpdateConn(ctx context.Context, conn *sqlx.DB, namespace *System) (*System, error) {
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
		return nil, err
	}

	var g System
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) Insert(ctx context.Context, namespace *System) (*System, error) {
	return InsertConn(ctx, r.conn, namespace)
}

func InsertConn(ctx context.Context, conn *sqlx.DB, namespace *System) (*System, error) {
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

	rows, err := conn.NamedQueryContext(ctx, query, namespace)
	if err != nil {
		return nil, err
	}

	var g System
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}

func (r *repository) GetById(ctx context.Context, id int) (*System, error) {
	return GetByIdConn(ctx, r.conn, id)
}

func GetByIdConn(ctx context.Context, conn *sqlx.DB, id int) (*System, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from systems
	   where id = $1
`
	var group System
	err := conn.GetContext(ctx, &group, query, id)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return DeleteConn(ctx, r.conn, id)
}

func  DeleteConn(ctx context.Context, conn *sqlx.DB, id int) error {
	query := `
		update systems 
			set deleted_at = now()
		where id = $1;
`
	_, err := conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}