package logger

import (
	"context"
	"go.uber.org/zap"
	"net/http"
)

type loggerCtx struct{}

func New(env string) (*zap.SugaredLogger, error) {

	log, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	if env != "dev" {
		log, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	}

	return log.Sugar(), nil
}

func Inject(log *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loggerCtx{}, log)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
func GetFromContext(ctx context.Context) *zap.SugaredLogger {
	log, ok := ctx.Value(loggerCtx{}).(*zap.SugaredLogger)
	if ok {
		return log
	}
	// todo: add check in
	return nil
}
