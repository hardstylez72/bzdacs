package acs

import (
	"context"
	"github.com/hardstylez72/bzdacs/pkg/relations/grouproute"
	"github.com/hardstylez72/bzdacs/pkg/relations/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/relations/userroute"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) *repository {
	return &repository{conn: conn}
}

func (r *repository) GetUserRoutes(ctx context.Context, userId, namespaceId int) ([]UserRoute, error) {

	groupsRoutes, err := r.getRoutesFromGroups(ctx, userId, namespaceId)
	if err != nil {
		return nil, err
	}
	userRoutes, err := r.getRoutesFromUserRoutes(ctx, userId, namespaceId)
	if err != nil {
		return nil, err
	}

	return mergeRoutes(groupsRoutes, userRoutes), nil
}

func (r *repository) getRoutesFromGroups(ctx context.Context, userId, namespaceId int) ([]UserRoute, error) {
	routes := make([]UserRoute, 0)
	groups, _, err := usergroup.List(ctx, r.conn, usergroup.Filter{
		Page:         1,
		PageSize:     usergroup.NoPageSizeLimit,
		Pattern:      usergroup.NoPattern,
		BelongToUser: true,
		NamespaceId:  namespaceId,
		UserId:       userId,
	})
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		groupRoutes, _, err := grouproute.List(ctx, r.conn, grouproute.Filter{
			Page:          1,
			PageSize:      grouproute.NoPageSizeLimit,
			RoutePattern:  grouproute.NoPattern,
			BelongToGroup: true,
			NamespaceId:   group.NamespaceId,
			GroupId:       group.Id,
		})
		if err != nil {
			return nil, err
		}

		for i := range groupRoutes {
			routes = append(routes, UserRoute{
				Route: Route{
					Id:          groupRoutes[i].Id,
					Route:       groupRoutes[i].Route,
					Method:      groupRoutes[i].Method,
					Description: groupRoutes[i].Description,
					NamespaceId: groupRoutes[i].Id,
					Times:       groupRoutes[i].Times,
				},
				IsExcluded: false,
			})
		}
	}
	return routes, err
}

func (r *repository) getRoutesFromUserRoutes(ctx context.Context, userId, namespaceId int) ([]UserRoute, error) {
	routes := make([]UserRoute, 0)
	userRoutes, _, err := userroute.ListBelongUser(ctx, r.conn, userroute.Filter{
		Page:         1,
		PageSize:     userroute.NoPageSizeLimit,
		Pattern:      userroute.NoPattern,
		BelongToUser: true,
		NamespaceId:  namespaceId,
		UserId:       userId,
	})
	if err != nil {
		return nil, err
	}
	for i := range userRoutes {
		routes = append(routes, UserRoute{
			Route: Route{
				Id:          userRoutes[i].Id,
				Route:       userRoutes[i].Route.Route,
				Method:      userRoutes[i].Method,
				Description: userRoutes[i].Description,
				NamespaceId: userRoutes[i].Id,
				Times:       userRoutes[i].Times,
			},
			IsExcluded: userRoutes[i].IsExcluded,
		})
	}
	return routes, err
}

func mergeRoutes(groupRoutes, userRoutes []UserRoute) []UserRoute {
	m := make(map[int]UserRoute)

	for i := range groupRoutes {
		id := groupRoutes[i].Id
		_, exist := m[id]
		if !exist {
			m[id] = groupRoutes[i]
		}
	}

	for i := range userRoutes {
		id := userRoutes[i].Id
		_, exist := m[id]
		if !exist {
			m[id] = userRoutes[i]
		} else {
			m[id] = userRoutes[i]
		}
	}

	routes := make([]UserRoute, 0, len(m))
	for _, route := range m {
		routes = append(routes, route)
	}
	return routes
}
