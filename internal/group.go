package internal

import (
	"context"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/tests/group"
	"github.com/hardstylez72/bzdacs/tests/relations/grouproute"
	"github.com/hardstylez72/bzdacs/tests/route"
)

func resolveGroup(ctx context.Context, c *client.BZDACS, code, description string, namespaceId int64) (*group.Group, error) {
	g, err := group.GetByParams(ctx, c, &group.Group{Code: code, Description: description, NamespaceId: namespaceId})
	if err != nil {
		if err == group.ErrNotFound {
			g, err = group.Create(ctx, c, &group.Group{Code: code, Description: description, NamespaceId: namespaceId})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return g, nil
}

func resolveGroupAndRoutes(ctx context.Context, c *client.BZDACS, g *group.Group, routes []route.Route) error {

	currentRoutes, err := grouproute.List(ctx, c, g.NamespaceId, g.Id)
	if err != nil {
		return err
	}

	pairs := make([]grouproute.Pair, 0)

	for _, r := range routes {
		if !contains(currentRoutes, r.Route) {
			pairs = append(pairs, grouproute.Pair{
				GroupID: g.Id,
				RouteID: r.Id,
			})
		}
	}

	err = grouproute.Create(ctx, c, pairs)
	if err != nil {
		return err
	}
	return nil
}

func contains(routes []string, target string) bool {
	for _, r := range routes {
		if r == target {
			return true
		}
	}
	return false
}
