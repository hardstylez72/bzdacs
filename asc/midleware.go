package acsmw

import (
	"context"
	"net/http"
)

type GetParamsFunc func(req *http.Request) *InputParams
type InputParams struct {
	Login       string
	Route       string
	Namespace   string
	System      string
	RouteMethod string
}

type Service interface {
	CheckAccess(ctx context.Context, params *InputParams) (*CheckAccessResponse, error)
}

func AccessCheck(host string, fn GetParamsFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		s := NewService(host)

		fn := func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			params := fn(r)

			res, err := s.CheckAccess(ctx, params)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if !res.AccessAllowed {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
