package namespace

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

func (r *repository) Update(ctx context.Context, namespace *Namespace) (*Namespace, error) {
	return UpdateConn(ctx, r.conn, namespace)
}

func UpdateConn(ctx context.Context, conn storage.SqlDriver, namespace *Namespace) (*Namespace, error) {
	query := `
	update tags
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

	var g Namespace
	for rows.Next() {
		err = rows.StructScan(&g)
		if err != nil {
			return nil, err
		}
	}

	return &g, nil
}


func (r *repository) GetById(ctx context.Context, id int) (*Namespace, error) {


	return GetByIdConn(ctx, r.conn, id)
}

func GetByIdConn(ctx context.Context, conn storage.SqlDriver, id int) (*Namespace, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from tags
	   where id = $1
`
	var group Namespace
	err := conn.GetContext(ctx, &group, query, id)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	return DeleteConn(ctx, r.conn, id)
}

func  DeleteConn(ctx context.Context, conn storage.SqlDriver, id int) error {
	query := `
		update tags 
			set deleted_at = now()
		where id = $1;
`
	_, err := conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}