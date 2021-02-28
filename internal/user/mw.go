package user

import (
	"context"
	"github.com/go-openapi/runtime/security"
	"net/http"
)

type login struct {
}

func AuthFunc(ctx context.Context, login string, password string) (context.Context, interface{}, error) {
	println(1)
	// todo: add auth
	return ctx, login, nil
}

func Auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			var login string

			ba := security.BasicAuthCtx(AuthFunc)
			ok, userLogin, err := ba.Authenticate(r)
			if err != nil {
				login, err = ExtractLogin(r)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}
			if !ok {
				login, err = ExtractLogin(r)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}
			if login == "" {
				login = userLogin.(string)
			}

			ctx = setLoginIntoContext(ctx, login)
			next.ServeHTTP(w, r.WithContext(ctx))
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
