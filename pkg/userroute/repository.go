package userroute

import (
	"context"
	"github.com/hardstylez72/bblog/ad/pkg/group"
	"github.com/hardstylez72/bblog/ad/pkg/grouproute"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) deletePair(ctx context.Context, tx *sqlx.Tx, routeId, userId int) error {
	query := `delete from ad.users_routes where user_id = $1 and route_id = $2`

	_, err := tx.ExecContext(ctx, query, userId, routeId)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, params []params) error {
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

func (r *repository) Update(ctx context.Context, params params) (*Route, error) {

	query := `
	with insert_row as (
		update ad.users_routes 
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
		from ad.routes r where r.id = $3;`

	rows := r.conn.QueryRowxContext(ctx, query, params.IsExcluded, params.UserId, params.RouteId)

	var route Route
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func insertPair(ctx context.Context, tx *sqlx.Tx, routeId, userId int, isExcluded bool) (*Route, error) {
	query := `
		with insert_row as (
			insert into ad.users_routes (
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
		from ad.routes r where r.id = $2;
`

	rows := tx.QueryRowxContext(ctx, query, userId, routeId, isExcluded)
	var route Route
	err := rows.StructScan(&route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func (r *repository) Insert(ctx context.Context, params []params) ([]Route, error) {

	tx, err := r.conn.BeginTxx(ctx, nil)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	routes := make([]Route, 0)
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
			Route:         routes[i],
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
		groupIds, err := grouproute.GetGroupIdsByRouteIdDb(ctx, r.conn, routes[i].Id)
		if err != nil {
			return nil, err
		}

		groups := make([]group.Group, 0)
		for j := range groupIds {
			gr, err := group.GetByIdDb(ctx, r.conn, groupIds[j])
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
			Route:         groupRoutes[i],
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
			Route:         overwrittenRoutes[i],
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

func ListGroupRoutesBelongUserDb(ctx context.Context, conn *sqlx.DB, userId int) ([]Route, error) {
	query := `
select r.id,
			   r.route,
               r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at,
		       false as is_excluded
		from ad.routes r
   inner join ad.groups_routes gr on gr.route_id = r.id
   inner join ad.groups g on gr.group_id = g.id
       where r.deleted_at is null
         and gr.group_id in (select group_id from ad.users_groups where user_id = $1)
         and g.deleted_at is null
    group by r.id
`
	routes := make([]Route, 0)
	err := conn.SelectContext(ctx, &routes, query, userId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func ListOverwrittenRoutesBelongUserDb(ctx context.Context, conn *sqlx.DB, userId int) ([]Route, error) {
	query := `
		select rg.route_id as id,
			   r.route,
               r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at,
		       rg.is_excluded
		from ad.routes r
    left join ad.users_routes rg on rg.route_id = r.id
        where rg.user_id = $1
          and deleted_at is null
`
	routes := make([]Route, 0)
	err := conn.SelectContext(ctx, &routes, query, userId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func ListRoutesNotBelongUserDb(ctx context.Context, conn *sqlx.DB, userId int) ([]Route, error) {
	query := `
	    select r.id,
			   r.route,
               r.method,
			   r.description,
			   r.created_at,
			   r.updated_at,
			   r.deleted_at
		from ad.routes r

       where r.deleted_at is null
         and r.id not in (
                       select rg.route_id as id
                         from ad.routes r
                   inner join ad.users_routes rg on rg.route_id = r.id
                        where rg.user_id = $1
                          and deleted_at is null
                union
                      select r.id
                        from ad.routes r
                  inner join ad.groups_routes gr on gr.route_id = r.id
                       where r.deleted_at is null
                         and gr.group_id in (select group_id from ad.users_groups where user_id = $1)
                    group by r.id
           )
`
	routes := make([]Route, 0)
	err := conn.SelectContext(ctx, &routes, query, userId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}
