package user

import (
	"context"
	"github.com/go-openapi/runtime/security"
	"github.com/hardstylez72/bzdacs/internal/session"
	"github.com/hardstylez72/bzdacs/pkg/util"
	"net/http"
)

type login struct {
}

func Auth(service *session.Service, user Repository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			var login string

			session, err := getSession(r, service)
			if err == nil {
				login = session.Login
				ctx = setLoginIntoContext(ctx, login)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			authFn := func(ctx context.Context, login string, password string) (context.Context, interface{}, error) {
				u, err := user.GetByParams(ctx, login, Encode(password))
				if err != nil {
					return nil, nil, err
				}
				return ctx, u.Login, nil
			}

			ok, userLogin, err := security.BasicAuthCtx(authFn).Authenticate(r)
			if err == nil && ok {
				if login == "" {
					login, ok = userLogin.(string)
				}
				if !ok {
					util.NewResp(w, r).Status(http.StatusUnauthorized).Send()
					return
				}
				ctx = setLoginIntoContext(ctx, login)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			util.NewResp(w, r).Status(http.StatusUnauthorized).Send()
		}

		return http.HandlerFunc(fn)
	}
}

func setLoginIntoContext(ctx context.Context, groups string) context.Context {
	return context.WithValue(ctx, login{}, groups)
}

func GetLoginFromContext(ctx context.Context) string {
	u, ok := ctx.Value(login{}).(string)
	if ok {
		return u
	}
	return ""
}
