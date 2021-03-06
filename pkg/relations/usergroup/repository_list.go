package usergroup

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

const NoPageSizeLimit = 0
const NoPattern = ""

type groupWithTotal struct {
	Group
	Total int `db:"total"`
}

func (r *repository) List(ctx context.Context, filter Filter) ([]Group, int, error) {

	routes, total, err := List(ctx, r.conn, filter)
	if err != nil {
		return nil, 0, err
	}

	return routes, total, nil
}

func List(ctx context.Context, driver storage.SqlDriver, f Filter) ([]Group, int, error) {

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Select(`
			   g.id,
			   g.code,
			   g.description,
			   g.created_at,
			   g.updated_at,
			   g.deleted_at,
			   g.namespace_id,
		       count(*) over() as total
			`).From("groups g").
		Where(sq.Eq{"g.deleted_at": nil}).
		Where(sq.Eq{"g.namespace_id": f.NamespaceId}).
		Offset(uint64((f.Page - 1) * f.PageSize)).
		OrderBy("g.updated_at desc")

	if f.BelongToUser {
		builder = builder.Join(`users_groups ug 
												on g.id = ug.group_id 
											   and ug.user_id = ?
											   and g.namespace_id = ?`, f.UserId, f.NamespaceId)
	} else {
		builder = builder.Where(`g.id not in (select group_id from users_groups where user_id = ?)`, f.UserId)
	}

	if f.PageSize != 0 {
		builder = builder.Limit(uint64(f.PageSize))
	}

	if f.Pattern != "" {
		f.Pattern = "%" + f.Pattern + "%"
		builder = builder.Where(sq.Like{"g.code": f.Pattern})
	}
	total := 0
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, total, storage.PgError(err)
	}
	groupsWithTotal := make([]groupWithTotal, 0)
	err = driver.SelectContext(ctx, &groupsWithTotal, query, args...)
	if err != nil {
		return nil, total, storage.PgError(err)
	}

	groups := make([]Group, 0, len(groupsWithTotal))
	for i, g := range groupsWithTotal {
		if i == 0 {
			total = g.Total
		}
		groups = append(groups, g.Group)
	}

	return groups, total, nil
}
