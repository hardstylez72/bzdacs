package route

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/relations/routetag"
	"github.com/hardstylez72/bzdacs/pkg/tag"
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
	newRoute, err := Insert(ctx, txx, route)
	if err != nil {
		return nil, err
	}

	tags, err := routetag.Merge(ctx, txx, newRoute.Id, route.Tags, route.NamespaceId)
	if err != nil {
		return nil, err
	}

	newRoute.Tags = getTagIdsFromArray(tags)

	return newRoute, nil
}

func getTagIdsFromArray(tags []tag.Tag) []string {
	names := make([]string, 0, len(tags))

	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	return names
}

func Insert(ctx context.Context, driver storage.SqlDriver, route *Route) (*Route, error) {
	query := `
insert into routes (
                       route,
                       method,
                       description,
                       created_at,
                       updated_at,
                       deleted_at,
					   namespace_id
                       )
                   values (
                       :route,
                       :method,
                       :description,
                       now(),
                       now(),
                       null,
					   :namespace_id
                   ) returning id,
                               route,
                       		   method,
                               description,
                               created_at,
                               updated_at,
                               deleted_at,
							   namespace_id
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
