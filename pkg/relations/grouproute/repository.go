package grouproute

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

var (
	ErrEntityNotFound = util.ErrEntityNotFound
)

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) deletePair(ctx context.Context, driver storage.SqlDriver, groupId, routeId int) error {
	query := `delete from groups_routes where route_id = $1 and group_id = $2`

	_, err := driver.ExecContext(ctx, query, routeId, groupId)
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
	txx := storage.WrapSqlxTx(tx)

	for _, pair := range params {

		err := r.deletePair(ctx, txx, pair.GroupId, pair.RouteId)
		if err != nil {
			return err
		}

	}
	return nil
}

func (r *repository) Insert(ctx context.Context, params []Pair) ([]Route, error) {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	txx := storage.WrapSqlxTx(tx)
	return Insert(ctx, txx, params)
}

func Insert(ctx context.Context, driver storage.SqlDriver, params []Pair) ([]Route, error) {

	routes := make([]Route, 0)
	for _, pair := range params {

		route, err := InsertPair(ctx, driver, pair.GroupId, pair.RouteId)
		if err != nil {
			return nil, err
		}
		routes = append(routes, *route)
	}

	return routes, nil
}

func InsertPair(ctx context.Context, driver storage.SqlDriver, groupId, routeId int) (*Route, error) {
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
			   r.deleted_at,
			   r.namespace_id
		from routes r where r.id = $1;
`

	rows := driver.QueryRowxContext(ctx, query, routeId, groupId)
	var route Route
	err := rows.StructScan(&route)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &route, nil
}

func GetGroupIdsByRouteId(ctx context.Context, driver storage.SqlDriver, routeId, namespaceId int) ([]int, error) {
	query := `select gr.group_id 
				from groups_routes gr
			    join groups g on gr.group_id = g.id and g.namespace_id = $2
			   where gr.route_id = $1
			   	`
	groupIds := make([]int, 0)
	err := driver.SelectContext(ctx, &groupIds, query, routeId, namespaceId)
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
