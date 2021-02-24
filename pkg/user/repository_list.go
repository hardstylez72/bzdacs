package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
)

type userWithTotal struct {
	User
	Total int `db:"total"`
}

func (u *repository) List(ctx context.Context, f filter) ([]User, int, error) {

	users, total, err := ListLL(ctx, u.conn, f)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func ListLL(ctx context.Context, driver storage.SqlDriver, f filter) ([]User, int, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rb := psql.Select(`
			   u.id,
			   u.external_id,
			   u.created_at,
			   u.updated_at,
			   u.deleted_at,
			   u.namespace_id,
			   count(*) over() as total
			`).
		From("users u").
		Where(sq.Eq{"u.deleted_at": nil}).
		Where(sq.Eq{"u.namespace_id": f.NamespaceId}).
		Offset(uint64((f.Page - 1) * f.PageSize)).
		OrderBy("u.updated_at desc")

	if f.PageSize != 0 {
		rb = rb.Limit(uint64(f.PageSize))
	}

	if f.Pattern != "" {
		f.Pattern = "%" + f.Pattern + "%"
		rb = rb.Where(sq.Like{"u.external_id": f.Pattern})
	}

	total := 0
	query, args, err := rb.ToSql()
	if err != nil {
		return nil, total, err
	}

	usersWithTotal := make([]userWithTotal, 0)
	err = driver.SelectContext(ctx, &usersWithTotal, query, args...)
	if err != nil {
		return nil, total, err
	}
	users := make([]User, 0, len(usersWithTotal))
	for i, u := range usersWithTotal {
		if i == 0 {
			total = u.Total
		}
		users = append(users, u.User)
	}

	return users, total, nil
}
