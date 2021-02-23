package usergroup

import (
	"context"
	"database/sql"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

var ErrEntityNotFound = util.ErrEntityNotFound

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) deletePair(ctx context.Context, tx *sqlx.Tx, groupId, userId int) error {
	query := `delete from users_groups where user_id = $1 and group_id = $2`

	_, err := tx.ExecContext(ctx, query, userId, groupId)
	if err != nil {
		return err
	}

	return nil
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

	for _, pair := range params {

		err := r.deletePair(ctx, tx, pair.GroupId, pair.UserId)
		if err != nil {
			return err
		}

	}
	return nil
}

func InsertPairTx(ctx context.Context, tx *sqlx.Tx, groupId, userId int) (*Group, error) {
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
			   r.deleted_at
		from groups r where r.id = $2;
`

	rows := tx.QueryRowxContext(ctx, query, userId, groupId)
	var route Group
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func InsertPairDb(ctx context.Context, conn *sqlx.DB, params Pair) (*Group, error) {
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
			   r.deleted_at
		from groups r where r.id = $2;
`

	rows := conn.QueryRowxContext(ctx, query, params.UserId, params.GroupId)
	var g Group
	err := rows.StructScan(&g)
	if err != nil {
		if g.Id == 0 {
			return nil, err
		}
		return nil, err
	}

	return &g, nil
}

func (r *repository) IsPairExist(ctx context.Context, pair Pair) (bool, error) {
	query := `select count(*) from users_groups where user_id = $1 and group_id = $2`

	var cnt = 0
	err := r.conn.GetContext(ctx, &cnt, query, pair.UserId, pair.GroupId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, ErrEntityNotFound
		}
		return false, err
	}

	return cnt >= 1, nil
}

func (r *repository) Insert(ctx context.Context, pair Pair) (*Group, error) {
	return InsertPairDb(ctx, r.conn, pair)
}

func (r *repository) InsertMany(ctx context.Context, params []Pair) ([]Group, error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	routes := make([]Group, 0)
	for _, pair := range params {

		route, err := InsertPairTx(ctx, tx, pair.GroupId, pair.UserId)
		if err != nil {
			return nil, err
		}
		routes = append(routes, *route)
	}

	return routes, nil
}

func (r *repository) ListUserNotInGroups(ctx context.Context, groupId int) ([]Group, error) {
	query := `
		select r.id,
			   r.code,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from groups r
        where r.id not in (select group_id from users_groups where user_id = $1)
          and deleted_at is null
`
	routes := make([]Group, 0)
	err := r.conn.SelectContext(ctx, &routes, query, groupId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func (r *repository) ListUserGroups(ctx context.Context, userId int) ([]Group, error) {
	query := `
		select rg.group_id as id,
			   r.code,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from groups r
    left join users_groups rg on rg.group_id = r.id
        where rg.user_id = $1 
          and deleted_at is null
`
	routes := make([]Group, 0)
	err := r.conn.SelectContext(ctx, &routes, query, userId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}
