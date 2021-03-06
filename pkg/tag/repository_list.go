package tag

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

type tagWithTotal struct {
	Tag
	Total int `db:"total"`
}

func (r *repository) List(ctx context.Context, filter filter) ([]Tag, int, error) {
	return List(ctx, r.conn, filter)
}

func List(ctx context.Context, driver storage.SqlDriver, f filter) ([]Tag, int, error) {

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.Select(`
 			   id,
			   name,
			   created_at,
			   updated_at,
			   deleted_at,
			   namespace_id,
		       count(*) over() as total
			`).From("tags").
		Where(sq.Eq{"deleted_at": nil}).
		Where(sq.Eq{"namespace_id": f.NamespaceId}).
		Limit(uint64(f.PageSize)).
		Offset(uint64((f.Page - 1) * f.PageSize)).
		OrderBy("updated_at desc")

	if f.Pattern != "" {
		f.Pattern = "%" + f.Pattern + "%"
		builder = builder.Where(sq.Like{"name": f.Pattern})
	}
	total := 0
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, total, storage.PgError(err)
	}
	tagsWithTotal := make([]tagWithTotal, 0)
	err = driver.SelectContext(ctx, &tagsWithTotal, query, args...)
	if err != nil {
		return nil, total, storage.PgError(err)
	}

	tags := make([]Tag, 0, len(tagsWithTotal))
	for i, r := range tagsWithTotal {
		if i == 0 {
			total = r.Total
		}
		tags = append(tags, r.Tag)
	}

	return tags, total, nil
}
