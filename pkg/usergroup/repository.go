package usergroup

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) deletePair(ctx context.Context, tx *sqlx.Tx, groupId, userId int) error {
	query := `delete from ad.users_groups where user_id = $1 and group_id = $2`

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

func (r *repository) insertPair(ctx context.Context, tx *sqlx.Tx, groupId, userId int) (*Group, error) {
	query := `
		with insert_row as (
			insert into ad.users_groups (
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
		from ad.groups r where r.id = $2;
`

	rows := tx.QueryRowxContext(ctx, query, userId, groupId)
	var route Group
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	return &route, nil
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

	routes := make([]Group, 0)
	for _, pair := range params {

		route, err := r.insertPair(ctx, tx, pair.GroupId, pair.UserId)
		if err != nil {
			return nil, err
		}
		routes = append(routes, *route)
	}

	return routes, nil
}

func (r *repository) ListNotInGroup(ctx context.Context, groupId int) ([]Group, error) {
	query := `
		select r.id,
			   r.code,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from ad.groups r
        where r.id not in (select group_id from ad.users_groups where user_id = $1)
          and deleted_at is null
`
	routes := make([]Group, 0)
	err := r.conn.SelectContext(ctx, &routes, query, groupId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func (r *repository) List(ctx context.Context, groupId int) ([]Group, error) {
	query := `
		select rg.group_id as id,
			   r.code,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from ad.groups r
    left join ad.users_groups rg on rg.group_id = r.id
        where rg.user_id = $1 
          and deleted_at is null
`
	routes := make([]Group, 0)
	err := r.conn.SelectContext(ctx, &routes, query, groupId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}
