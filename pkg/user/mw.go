package user

import (
	"context"
	"net/http"
)

type login struct {
}

func Auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		fn := func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			login, err := ExtractLogin(r)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
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
