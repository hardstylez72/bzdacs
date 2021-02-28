package internal

import (
	"context"
	"github.com/hardstylez72/bzdacs/generated/client"
	rt "github.com/hardstylez72/bzdacs/pkg/route"
	"github.com/hardstylez72/bzdacs/tests/route"

	"strings"
)

func resolveRoutes(ctx context.Context, c *client.BZDACS, routes []rt.Route, namespaceId int64) ([]route.Route, error) {
	rs := make([]route.Route, 0)
	for _, r := range routes {
		r, err := resolveRoute(ctx, c, r, namespaceId)
		if err != nil {
			return nil, err
		}
		rs = append(rs, *r)
	}
	return rs, nil
}

func resolveRoute(ctx context.Context, c *client.BZDACS, r rt.Route, namespaceId int64) (*route.Route, error) {
	s, err := route.GetByParams(ctx, c, &route.Route{
		Route:       r.Route,
		Method:      r.Method,
		NamespaceId: namespaceId,
	})
	if err != nil {
		if err == route.ErrNotFound {
			s, err = route.Create(ctx, c, &route.Route{
				Route:       r.Route,
				Method:      r.Method,
				Description: r.Description,
				NamespaceId: namespaceId,
				Tags:        r.Tags,
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}

	}
	return s, nil
}

func filterGuestRoutes(in []route.Route) []route.Route {
	out := make([]route.Route, 0)

	for _, r := range in {
		if strings.Contains(r.Route, "/update") ||
			strings.Contains(r.Route, "/delete") ||
			strings.Contains(r.Route, "/create") {
			continue
		}
		out = append(out, r)
	}
	return out
}
