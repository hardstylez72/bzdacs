package userroute

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/relations/grouproute"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) Delete(ctx context.Context, params []PairToDelete) error {
	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	for _, pair := range params {

		err := r.deletePairLL(ctx, tx, pair.RouteId, pair.UserId)
		if err != nil {
			return err
		}

	}
	return nil
}

func (r *repository) deletePairLL(ctx context.Context, tx *sqlx.Tx, routeId, userId int) error {
	query := `delete from users_routes where user_id = $1 and route_id = $2`

	_, err := tx.ExecContext(ctx, query, userId, routeId)
	if err != nil {
		return storage.PgError(err)
	}

	return nil
}

func (r *repository) Update(ctx context.Context, params Pair) (*Route, error) {
	return UpdateLL(ctx, r.conn, params)
}
func UpdateLL(ctx context.Context, driver storage.SqlDriver, params Pair) (*Route, error) {

	query := `
	with insert_row as (
		update users_routes 
		   set is_excluded = $1 
		 where user_id = $2 and route_id = $3
	)
		select r.id,
			   r.route,
               r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at,
		       $1 as is_excluded
		from routes r where r.id = $3;`

	rows := driver.QueryRowxContext(ctx, query, params.IsExcluded, params.UserId, params.RouteId)
	var route Route
	err := rows.StructScan(&route)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &route, nil
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

	routes := make([]Route, 0)
	for _, pair := range params {

		route, err := insertPairLL(ctx, txx, pair.RouteId, pair.UserId, pair.IsExcluded)
		if err != nil {
			return nil, err
		}
		routes = append(routes, *route)
	}

	return routes, nil
}

func insertPairLL(ctx context.Context, driver storage.SqlDriver, routeId, userId int, isExcluded bool) (*Route, error) {
	query := `
		with insert_row as (
			insert into users_routes (
					   user_id,
					   route_id,
			           is_excluded
					   )
				   values (
					   $1,
					   $2,
				       $3
				   )
		)
		select r.id,
			   r.route,
               r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at,
		       $3 as is_excluded
		from routes r where r.id = $2;
`

	rows := driver.QueryRowxContext(ctx, query, userId, routeId, isExcluded)
	var route Route
	err := rows.StructScan(&route)
	if err != nil {
		return nil, storage.PgError(err)
	}

	return &route, nil
}

func (r *repository) getRouteGroups(ctx context.Context, routeId, namespaceId int) ([]Group, error) {

	groupIds, err := grouproute.GetGroupIdsByRouteIdLL(ctx, r.conn, routeId, namespaceId)
	if err != nil {
		return nil, err
	}

	groups := make([]Group, 0)
	for j := range groupIds {
		g, err := group.GetByIdLL(ctx, r.conn, groupIds[j], namespaceId)
		if err != nil {
			return nil, err
		}
		if g.DeletedAt.Valid {
			continue
		}

		groups = append(groups, Group{
			Id:          g.Id,
			Code:        g.Code,
			Description: g.Description,
			NamespaceId: g.NamespaceId,
			Times: util.Times{
				CreatedAt: g.CreatedAt,
				UpdatedAt: g.UpdatedAt,
				DeletedAt: g.DeletedAt,
			},
		})
	}

	return groups, nil
}
