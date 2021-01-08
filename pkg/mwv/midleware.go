package mwv

import (
	"context"
	"net/http"
)

type userGroups struct{}

type GetLoginFunc func(req *http.Request) string
type GetRouteAndMethodFunc func(req *http.Request) (route, method string)

type Service interface {
	CheckAccess(ctx context.Context, login, route, method string) (*CheckAccessResponse, error)
}

func AccessCheck(service Service, getLoginFn GetLoginFunc, getRouteAndMethodFn GetRouteAndMethodFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			login := getLoginFn(r)
			route, method := getRouteAndMethodFn(r)

			res, err := service.CheckAccess(ctx, login, route, method)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if !res.AccessAllowed {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			ctx = setUserGroupsIntoContext(ctx, res.Groups)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

func setUserGroupsIntoContext(ctx context.Context, groups []string) context.Context {
	return context.WithValue(ctx, userGroups{}, groups)
}

func GetUserGroupsFromContext(ctx context.Context) []string {
	u, ok := ctx.Value(userGroups{}).([]string)
	if ok {
		return u
	}
	return []string{""}
}
