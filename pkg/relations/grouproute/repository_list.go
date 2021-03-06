package grouproute

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/route"
)

const NoPageSizeLimit = 0
const NoPattern = ""

type routeWithTotal struct {
	Route
	Total int `db:"total"`
}

func (r *repository) List(ctx context.Context, filter Filter) ([]Route, int, error) {

	routes, total, err := List(ctx, r.conn, filter)
	if err != nil {
		return nil, 0, err
	}

	for i := range routes {
		tagNames, err := route.GetTagNamesByRouteId(ctx, r.conn, routes[i].Id, routes[i].NamespaceId)
		if err != nil {
			return nil, 0, err
		}

		routes[i].Tags = tagNames
	}

	return routes, total, nil
}

func List(ctx context.Context, driver storage.SqlDriver, f Filter) ([]Route, int, error) {

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Select(`
			   r.id,
			   r.route,
			   r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at,
			   r.namespace_id,
		       count(*) over() as total
			`).From("routes r").
		Where(sq.Eq{"r.deleted_at": nil}).
		Where(sq.Eq{"r.namespace_id": f.NamespaceId}).
		Offset(uint64((f.Page - 1) * f.PageSize)).
		OrderBy("r.updated_at desc")

	if f.BelongToGroup {
		builder = builder.Join(`groups_routes gr 
												on r.id = gr.route_id 
											   and gr.group_id = ?
											   and r.namespace_id = ?`, f.GroupId, f.NamespaceId)
	} else {
		builder = builder.Where(`r.id not in (select route_id from groups_routes where group_id = ?)`, f.GroupId)
	}

	if f.PageSize != 0 {
		builder = builder.Limit(uint64(f.PageSize))
	}

	if f.RoutePattern != "" {
		f.RoutePattern = "%" + f.RoutePattern + "%"
		builder = builder.Where(sq.Like{"r.route": f.RoutePattern})
	}
	total := 0
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, total, storage.PgError(err)
	}
	tagsWithTotal := make([]routeWithTotal, 0)
	err = driver.SelectContext(ctx, &tagsWithTotal, query, args...)
	if err != nil {
		return nil, total, storage.PgError(err)
	}

	tags := make([]Route, 0, len(tagsWithTotal))
	for i, r := range tagsWithTotal {
		if i == 0 {
			total = r.Total
		}
		tags = append(tags, r.Route)
	}

	return tags, total, nil
}
