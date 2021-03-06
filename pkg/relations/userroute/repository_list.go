package userroute

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

const NoPageSizeLimit = 0
const NoPattern = ""

type RouteWithGroupsWithTotal struct {
	RouteWithGroups
	Total int `db:"total"`
}

func (r *repository) List(ctx context.Context, filter Filter) (routes []RouteWithGroups, total int, err error) {

	if filter.BelongToUser {
		routes, total, err = ListBelongUser(ctx, r.conn, filter)
	} else {
		routes, total, err = ListNotBelongUser(ctx, r.conn, filter)
	}

	if err != nil {
		return nil, 0, err
	}

	for i := range routes {
		groups, err := r.getRouteGroups(ctx, routes[i].Id, routes[i].NamespaceId)
		if err != nil {
			return nil, -1, err
		}
		routes[i].Groups = groups
	}

	return routes, total, nil
}

func ListNotBelongUser(ctx context.Context, driver storage.SqlDriver, f Filter) ([]RouteWithGroups, int, error) {

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
		       false as is_excluded,
			   count(*) over() as total
			`).
		From("routes r").
		Where("r.id not in (select route_id from users_routes where user_id = ?)", f.UserId).
		Where(sq.Eq{"r.deleted_at": nil}).
		Where(sq.Eq{"r.namespace_id": f.NamespaceId}).
		Offset(uint64((f.Page - 1) * f.PageSize)).
		OrderBy("r.updated_at desc")

	if f.PageSize != 0 {
		builder = builder.Limit(uint64(f.PageSize))
	}

	if f.Pattern != "" {
		f.Pattern = "%" + f.Pattern + "%"
		builder = builder.Where(sq.Like{"r.route": f.Pattern})
	}
	total := 0
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, total, storage.PgError(err)
	}
	RouteWithGroupsWithTotal := make([]RouteWithGroupsWithTotal, 0)
	err = driver.SelectContext(ctx, &RouteWithGroupsWithTotal, query, args...)
	if err != nil {
		return nil, total, storage.PgError(err)
	}

	RouteWithGroups := make([]RouteWithGroups, 0, len(RouteWithGroupsWithTotal))
	for i, g := range RouteWithGroupsWithTotal {
		if i == 0 {
			total = g.Total
		}
		RouteWithGroups = append(RouteWithGroups, g.RouteWithGroups)
	}

	return RouteWithGroups, total, nil
}
func ListBelongUser(ctx context.Context, driver storage.SqlDriver, f Filter) ([]RouteWithGroups, int, error) {

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
		       rg.is_excluded,
			   count(*) over() as total
			`).
		From("routes r").
		Join("users_routes rg on rg.route_id = r.id and rg.user_id = ?", f.UserId).
		Where(sq.Eq{"r.deleted_at": nil}).
		Where(sq.Eq{"r.namespace_id": f.NamespaceId}).
		Offset(uint64((f.Page - 1) * f.PageSize)).
		OrderBy("r.updated_at desc")

	if f.PageSize != 0 {
		builder = builder.Limit(uint64(f.PageSize))
	}

	if f.Pattern != "" {
		f.Pattern = "%" + f.Pattern + "%"
		builder = builder.Where(sq.Like{"r.route": f.Pattern})
	}
	total := 0
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, total, storage.PgError(err)
	}
	RouteWithGroupsWithTotal := make([]RouteWithGroupsWithTotal, 0)
	err = driver.SelectContext(ctx, &RouteWithGroupsWithTotal, query, args...)
	if err != nil {
		return nil, total, storage.PgError(err)
	}

	RouteWithGroups := make([]RouteWithGroups, 0, len(RouteWithGroupsWithTotal))
	for i, g := range RouteWithGroupsWithTotal {
		if i == 0 {
			total = g.Total
		}
		RouteWithGroups = append(RouteWithGroups, g.RouteWithGroups)
	}

	return RouteWithGroups, total, nil
}
