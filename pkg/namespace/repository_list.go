package namespace

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

const (
	ListBySystemId ListKind = "ListBySystemId"
)

func (r *repository) List(ctx context.Context, kind ListKind, systemId int, arg ...interface{}) ([]Namespace, error) {
	switch kind {
	case ListBySystemId:
		return List(ctx, r.conn, systemId)
	}
	return nil, errors.New("invalid input params")
}

func List(ctx context.Context, conn storage.SqlDriver, systemId int) ([]Namespace, error) {
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
