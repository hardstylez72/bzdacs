package grouproute

import (
	"context"
	"database/sql"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

var (
	ErrEntityAlreadyExists = util.ErrEntityAlreadyExists
	ErrEntityNotFound      = util.ErrEntityNotFound
)

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) deletePair(ctx context.Context, tx *sqlx.Tx, groupId, routeId int) error {
	query := `delete from groups_routes where route_id = $1 and group_id = $2`

	_, err := tx.ExecContext(ctx, query, routeId, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Insert(ctx context.Context, params Pair) (*Route, error) {
	return InsertPairDb(ctx, r.conn, params.GroupId, params.RouteId)
}

func (r *repository) IsPairExist(ctx context.Context, pair Pair) (bool, error) {
	query := `select count(*) from groups_routes where route_id = $1 and group_id = $2`

	var cnt = 0
	err := r.conn.GetContext(ctx, &cnt, query, pair.RouteId, pair.GroupId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, ErrEntityNotFound
		}
		return false, err
	}

	return cnt >= 1, nil
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

		err := r.deletePair(ctx, tx, pair.GroupId, pair.RouteId)
		if err != nil {
			return err
		}

	}
	return nil
}

func (r *repository) InsertMany(ctx context.Context, params []Pair) ([]Route, error) {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	return InsertTx(ctx, tx, params)
}

func InsertTx(ctx context.Context, tx *sqlx.Tx, params []Pair) ([]Route, error) {

	routes := make([]Route, 0)
	for _, pair := range params {

		route, err := InsertPairTx(ctx, tx, pair.GroupId, pair.RouteId)
		if err != nil {
			return nil, err
		}
		routes = append(routes, *route)
	}

	return routes, nil
}

func InsertPairDb(ctx context.Context, conn *sqlx.DB, groupId, routeId int) (*Route, error) {
	query := `
		with insert_row as (
			insert into groups_routes (
					   route_id,
					   group_id
					   )
				   values (
					   $1,
					   $2
				   ) on conflict do nothing 
		)
		select r.id,
			   r.route,
		       r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from routes r where r.id = $1;
`

	rows := conn.QueryRowxContext(ctx, query, routeId, groupId)
	var route Route
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	if route.Id == 0 {
		return nil, util.ErrEntityAlreadyExists
	}

	return &route, nil
}

func InsertPairTx(ctx context.Context, tx *sqlx.Tx, groupId, routeId int) (*Route, error) {
	query := `
		with insert_row as (
			insert into groups_routes (
					   route_id,
					   group_id
					   )
				   values (
					   $1,
					   $2
				   ) on conflict do nothing 
		)
		select r.id,
			   r.route,
		       r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from routes r where r.id = $1;
`

	rows := tx.QueryRowxContext(ctx, query, routeId, groupId)
	var route Route
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	if route.Id == 0 {
		return nil, util.ErrEntityAlreadyExists
	}

	return &route, nil
}

func GetGroupIdsByRouteIdDb(ctx context.Context, conn *sqlx.DB, routeId int) ([]int, error) {
	query := `select gr.group_id from groups_routes gr where gr.route_id = $1`
	groupIds := make([]int, 0)
	err := conn.SelectContext(ctx, &groupIds, query, routeId)
	if err != nil {
		return nil, err
	}

	return groupIds, nil
}

func (r *repository) ListNotInGroup(ctx context.Context, groupId int) ([]Route, error) {
	// todo: fix add rg.is_excluded
	query := `
		select r.id,
			   r.route,
		       r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from routes r
        where r.id not in (select route_id from groups_routes where group_id = $1)
          and deleted_at is null
`
	routes := make([]Route, 0)
	err := r.conn.SelectContext(ctx, &routes, query, groupId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func (r *repository) List(ctx context.Context, groupId int) ([]Route, error) {
	return ListDb(ctx, r.conn, groupId)
}

func ListDb(ctx context.Context, conn *sqlx.DB, groupId int) ([]Route, error) {
	query := `
		select rg.route_id as id,
			   r.route,
		       r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from routes r
    left join groups_routes rg on rg.route_id = r.id
        where rg.group_id = $1 
          and deleted_at is null
`
	routes := make([]Route, 0)
	err := conn.SelectContext(ctx, &routes, query, groupId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}
