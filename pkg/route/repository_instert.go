package route

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/routetag"
)

func (r *repository) Insert(ctx context.Context, route *Route) (*Route, error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	txx := storage.WrapSqlxTx(tx)
	newRoute, err := InsertLL(ctx, txx, route)
	if err != nil {
		return nil, err
	}

	tagNames, err := routetag.Merge(ctx, r.conn, tx, newRoute.Id, route.Tags)
	if err != nil {
		return nil, err
	}
	newRoute.Tags = tagNames

	return newRoute, nil
}

func InsertLL(ctx context.Context, driver storage.SqlDriver, route *Route) (*Route, error) {
	query := `
insert into routes (
                       route,
                       method,
                       description,
                       created_at,
                       updated_at,
                       deleted_at
                       )
                   values (
                       :route,
                       :method,
                       :description,
                       now(),
                       now(),
                       null
                   ) returning id,
                               route,
                       		   method,
                               description,
                               created_at,
                               updated_at,
                               deleted_at;
`
	var r Route
	rows, err := driver.NamedQueryContext(ctx, query, route)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
	}

	return &r, nil
}
