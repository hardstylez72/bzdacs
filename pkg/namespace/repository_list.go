package namespace

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

func (r *repository) List(ctx context.Context, systemId int) ([]Namespace, error) {
	return ListLL(ctx, r.conn, systemId)
}

func ListLL(ctx context.Context, conn storage.SqlDriver, systemId int) ([]Namespace, error) {
	query := `
    select
           n.id,
           n.name,
           n.created_at,
           n.updated_at,
           n.deleted_at,
   		   n.system_id
      from namespaces n
     where n.system_id = $1
`
	ns := make([]Namespace, 0)
	err := conn.SelectContext(ctx, &ns, query, systemId)
	if err != nil {
		return nil, err
	}

	return ns, nil
}
