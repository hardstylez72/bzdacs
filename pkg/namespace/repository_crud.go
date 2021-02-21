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
	update namespaces
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
		from namespaces
	   where id = $1
`
	var namespace Namespace
	err := conn.GetContext(ctx, &namespace, query, id)
	if err != nil {
		return nil, err
	}

	return &namespace, nil
}

func (r *repository) Delete(ctx context.Context, systemId, namespaceId int) error {
	return DeleteLL(ctx, r.conn, systemId, namespaceId)
}

func DeleteLL(ctx context.Context, conn storage.SqlDriver, systemId, namespaceId int) error {
	query := `
		delete from systems_namespaces
		where system_id = $1 and namespace_id = $2
`
	_, err := conn.ExecContext(ctx, query, systemId, namespaceId)
	if err != nil {
		return err
	}

	return nil
}
