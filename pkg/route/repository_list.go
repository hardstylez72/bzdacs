package route

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

type routeWithTotal struct {
	Route
	Total int `db:"total"`
}

func (r *repository) List(ctx context.Context, f filter) ([]Route, int, error) {

	routes, total, err := ListLL(ctx, r.conn, f)
	if err != nil {
		return nil, 0, err
	}

	for i := range routes {

		tagNames, err := GetTagNamesByRouteId(ctx, r.conn, routes[i].Id, f.NamespaceId)
		if err != nil {
			return nil, 0, err
		}
		routes[i].Tags = tagNames

	}
	return routes, total, nil
}

func ListLL(ctx context.Context, driver storage.SqlDriver, f filter) ([]Route, int, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rb := psql.Select(`
			   r.id,
			   r.route,
		       r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at,
			   r.namespace_id,
			   count(*) over() as total
			`).
		From("routes r").
		Where(sq.Eq{"r.deleted_at": nil}).
		Where(sq.Eq{"r.namespace_id": f.NamespaceId}).
		Limit(uint64(f.PageSize)).
		Offset(uint64((f.Page - 1) * f.PageSize)).
		OrderBy("r.updated_at desc")

	total := 0
	query, args, err := rb.ToSql()
	if err != nil {
		return nil, total, err
	}

	routesWithTotal := make([]routeWithTotal, 0)
	err = driver.SelectContext(ctx, &routesWithTotal, query, args...)
	if err != nil {
		return nil, total, err
	}
	routes := make([]Route, 0, len(routesWithTotal))
	for i, r := range routesWithTotal {
		if i == 0 {
			total = r.Total
		}
		routes = append(routes, r.Route)
	}

	return routes, total, nil
}
