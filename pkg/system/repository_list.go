package system

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

const (
	ListNoFilter ListKind = "ListNoFilter"
)

func (r *repository) List(ctx context.Context, kind ListKind, arg ...interface{}) ([]System, error) {

	switch kind {
	case ListNoFilter:
		return List(ctx, r.conn)
	}
	return nil, errors.New("invalid input params")
}

func List(ctx context.Context, driver storage.SqlDriver) ([]System, error) {
	query := `
		select id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at
		from systems
	   where deleted_at is null;
`
	systems := make([]System, 0)
	err := driver.SelectContext(ctx, &systems, query)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return systems, nil
}
