package warmup

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/relations/grouproute"
	"github.com/hardstylez72/bzdacs/pkg/route"
)

func resolveGroup(ctx context.Context, driver storage.SqlDriver, code, description string, namespaceId int) (*group.Group, error) {
	g, err := group.GetByCode(ctx, driver, code, namespaceId)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			g, err = group.Insert(ctx, driver, &group.Group{Code: code, Description: description, NamespaceId: namespaceId})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return g, nil
}

func resolveGroupAndRoutes(ctx context.Context, driver storage.SqlDriver, g *group.Group, routes []route.Route) error {

	currentRoutes, _, err := grouproute.List(ctx, driver, grouproute.Filter{
		Page:          1,
		PageSize:      grouproute.NoPageSizeLimit,
		RoutePattern:  grouproute.NoPattern,
		BelongToGroup: true,
		NamespaceId:   g.NamespaceId,
		GroupId:       g.Id,
	})
	if err != nil {
		return err
	}

	pairs := make([]grouproute.Pair, 0)

	for _, r := range routes {
		if !containsRoutes(currentRoutes, r) {
			pairs = append(pairs, grouproute.Pair{
				GroupId: g.Id,
				RouteId: r.Id,
			})
		}
	}

	_, err = grouproute.Insert(ctx, driver, pairs)
	if err != nil {
		return err
	}
	return nil
}

func containsRoutes(routes []grouproute.Route, target route.Route) bool {
	for _, route := range routes {
		if route.Route == target.Route && route.Method == target.Method && route.NamespaceId == target.NamespaceId {
			return true
		}
	}
	return false
}
