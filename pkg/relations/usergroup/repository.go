package usergroup

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) Delete(ctx context.Context, params []Pair) error {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	txx := storage.WrapSqlxTx(tx)
	for _, pair := range params {
		err := r.deletePair(ctx, txx, pair.GroupId, pair.UserId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) deletePair(ctx context.Context, driver storage.SqlDriver, groupId, userId int) error {
	query := `delete from users_groups where user_id = $1 and group_id = $2`

	_, err := driver.ExecContext(ctx, query, userId, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Insert(ctx context.Context, params []Pair) ([]Group, error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	txx := storage.WrapSqlxTx(tx)

	routes := make([]Group, 0)
	for _, pair := range params {

		route, err := InsertPair(ctx, txx, pair.GroupId, pair.UserId)
		if err != nil {
			return nil, err
		}
		routes = append(routes, *route)
	}

	return routes, nil
}
func InsertPair(ctx context.Context, driver storage.SqlDriver, groupId, userId int) (*Group, error) {
	query := `
		with insert_row as (
			insert into users_groups (
					   user_id,
					   group_id
					   )
				   values (
					   $1,
					   $2
				   )
		)
		select r.id,
			   r.code,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at,
		       r.namespace_id
		from groups r where r.id = $2;
`

	rows := driver.QueryRowxContext(ctx, query, userId, groupId)
	var route Group
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}
