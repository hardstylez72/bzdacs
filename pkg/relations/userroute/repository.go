package userroute

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/relations/grouproute"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) deletePair(ctx context.Context, tx *sqlx.Tx, routeId, userId int) error {
	query := `delete from users_routes where user_id = $1 and route_id = $2`

	_, err := tx.ExecContext(ctx, query, userId, routeId)
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

		err := r.deletePair(ctx, tx, pair.RouteId, pair.UserId)
		if err != nil {
			return err
		}

	}
	return nil
}

func (r *repository) Update(ctx context.Context, params Pair) (*UserRoute, error) {

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

	rows := r.conn.QueryRowxContext(ctx, query, params.IsExcluded, params.UserId, params.RouteId)

	var route UserRoute
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func insertPair(ctx context.Context, tx *sqlx.Tx, routeId, userId int, isExcluded bool) (*UserRoute, error) {
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

	rows := tx.QueryRowxContext(ctx, query, userId, routeId, isExcluded)
	var route UserRoute
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func (r *repository) Insert(ctx context.Context, params []Pair) ([]UserRoute, error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	routes := make([]UserRoute, 0)
	for _, pair := range params {

		route, err := insertPair(ctx, tx, pair.RouteId, pair.UserId, pair.IsExcluded)
		if err != nil {
			return nil, err
		}
		routes = append(routes, *route)
	}

	return routes, nil
}

func (r *repository) RoutesNotBelongUser(ctx context.Context, userId int) ([]RouteWithGroups, error) {
	routes, err := ListRoutesNotBelongUserDb(ctx, r.conn, userId)
	if err != nil {
		return nil, err
	}

	routesExt := make([]RouteExt, 0)

	for i := range routes {
		routesExt = append(routesExt, RouteExt{

			Route:         routes[i].Route,
			IsOverwritten: false,
		})
	}

	routeWithGroups, err := r.enrichRoutesWithGroups(ctx, routesExt)
	if err != nil {
		return nil, err
	}

	return routeWithGroups, nil
}

func (r *repository) enrichRoutesWithGroups(ctx context.Context, routes []RouteExt) ([]RouteWithGroups, error) {
	routeWithGroups := make([]RouteWithGroups, 0)

	for i := range routes {
		groupIds, err := grouproute.GetGroupIdsByRouteIdLL(ctx, r.conn, routes[i].Id)
		if err != nil {
			return nil, err
		}

		groups := make([]group.Group, 0)
		for j := range groupIds {
			gr, err := group.GetByIdLL(ctx, r.conn, groupIds[j], routes[i].NamespaceId)
			if err != nil {
				return nil, err
			}
			if gr.DeletedAt.Valid {
				continue
			}

			groups = append(groups, *gr)
		}

		routeWithGroups = append(routeWithGroups, RouteWithGroups{
			Groups:   groups,
			RouteExt: routes[i],
		})
	}

	return routeWithGroups, nil
}
func (r *repository) RoutesBelongUser(ctx context.Context, userId int) ([]RouteWithGroups, error) {
	routes, err := ListRoutesBelongUserDb(ctx, r.conn, userId)
	if err != nil {
		return nil, err
	}

	routeWithGroups, err := r.enrichRoutesWithGroups(ctx, routes)
	if err != nil {
		return nil, err
	}

	return routeWithGroups, nil
}

func ListRoutesBelongUserDb(ctx context.Context, conn *sqlx.DB, userId int) ([]RouteExt, error) {
	groupRoutes, err := ListGroupRoutesBelongUserDb(ctx, conn, userId)
	if err != nil {
		return nil, err
	}
	overwrittenRoutes, err := ListOverwrittenRoutesBelongUserDb(ctx, conn, userId)
	if err != nil {
		return nil, err
	}
	routes := make([]RouteExt, 0)

	for i := range groupRoutes {
		routes = append(routes, RouteExt{
			Route:         groupRoutes[i].Route,
			IsOverwritten: false,
		})
	}

	for i := range overwrittenRoutes {
		contains, index := contains(routes, overwrittenRoutes[i].Id)
		if contains {
			routes[index].IsOverwritten = true
			routes[index].IsIndependent = true
			routes[index].IsExcluded = overwrittenRoutes[i].IsExcluded
			continue
		}

		routes = append(routes, RouteExt{
			Route:         overwrittenRoutes[i].Route,
			IsOverwritten: false,
			IsIndependent: true,
		})
	}

	return routes, nil
}

func contains(routes []RouteExt, id int) (bool, int) {
	for j := range routes {
		if id == routes[j].Id {
			return true, j
		}
	}
	return false, -1
}

func ListGroupRoutesBelongUserDb(ctx context.Context, conn *sqlx.DB, userId int) ([]UserRoute, error) {
	query := `
		select r.id,
			   r.route,
               r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at,
		       false as is_excluded
		from routes r
        join groups_routes gr on gr.route_id = r.id and gr.group_id in (select group_id from users_groups where user_id = $1)
        join groups g on gr.group_id = g.id and g.deleted_at is null
       where r.deleted_at is null
    group by r.id
`
	routes := make([]UserRoute, 0)
	err := conn.SelectContext(ctx, &routes, query, userId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func ListOverwrittenRoutesBelongUserDb(ctx context.Context, conn *sqlx.DB, userId int) ([]UserRoute, error) {
	query := `
		select rg.route_id as id,
			   r.route,
               r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at,
		       rg.is_excluded
		from routes r
        join users_routes rg on rg.route_id = r.id
        where rg.user_id = $1
          and deleted_at is null
`
	routes := make([]UserRoute, 0)
	err := conn.SelectContext(ctx, &routes, query, userId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func ListRoutesNotBelongUserDb(ctx context.Context, conn *sqlx.DB, userId int) ([]UserRoute, error) {
	query := `
	    select r.id,
			   r.route,
               r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from routes r

       where r.deleted_at is null
         and r.id not in (
                       select rg.route_id as id
                         from routes r
                   inner join users_routes rg on rg.route_id = r.id
                        where rg.user_id = $1
                          and deleted_at is null
                union
                      select r.id
                        from routes r
                  inner join groups_routes gr on gr.route_id = r.id
                       where r.deleted_at is null
                         and gr.group_id in (select group_id from users_groups where user_id = $1)
                    group by r.id
           )
`
	routes := make([]UserRoute, 0)
	err := conn.SelectContext(ctx, &routes, query, userId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}
