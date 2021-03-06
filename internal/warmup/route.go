package warmup

import (
	"context"
	"errors"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/route"
)

func resolveRoutes(ctx context.Context, driver storage.SqlDriver, routes []route.Route, namespaceId int) ([]route.Route, error) {
	rs := make([]route.Route, 0)
	for _, r := range routes {
		r, err := resolveRoute(ctx, driver, r, namespaceId)
		if err != nil {
			return nil, err
		}
		rs = append(rs, *r)
	}
	return rs, nil
}

func resolveRoute(ctx context.Context, driver storage.SqlDriver, r route.Route, namespaceId int) (*route.Route, error) {
	s, err := route.GetByParams(ctx, driver, r.Method, r.Route, namespaceId)
	if err != nil {
		if errors.Is(err, storage.EntityNotFound) {
			s, err = route.Insert(ctx, driver, &route.Route{
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
