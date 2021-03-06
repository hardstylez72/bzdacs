package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	acsmw "github.com/hardstylez72/bzdacs-go"
	"github.com/hardstylez72/bzdacs/internal/session"
	sysuser "github.com/hardstylez72/bzdacs/internal/user"
	"github.com/hardstylez72/bzdacs/internal/warmup"
	"github.com/hardstylez72/bzdacs/pkg/acs"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/infra/logger"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/namespace"
	"github.com/hardstylez72/bzdacs/pkg/relations/grouproute"
	"github.com/hardstylez72/bzdacs/pkg/relations/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/relations/userroute"
	"github.com/hardstylez72/bzdacs/pkg/route"
	"github.com/hardstylez72/bzdacs/pkg/system"
	"github.com/hardstylez72/bzdacs/pkg/tag"
	"github.com/hardstylez72/bzdacs/pkg/user"
	"github.com/spf13/viper"
	"github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

type Server struct {
	log    *zap.SugaredLogger
	router chi.Router

	repository struct {
		user       user.Repository
		tag        tag.Repository
		route      route.Repository
		group      group.Repository
		grouproute grouproute.Repository
		usergroup  usergroup.Repository
		userroute  userroute.Repository
		namespace  namespace.Repository
		system     system.Repository
		sysuser    sysuser.Repository
		acs        acs.Repository
	}
}

func (s *Server) startBackendServer(log *zap.SugaredLogger, done, ready chan struct{}) error {
	log.Info("app is initializing")

	r := chi.NewRouter()
	c := cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	r.Use(c)
	r.Use(logger.Inject(log))

	err := s.Start(r)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = warmup.Init(ctx, buildRoutes(r))
	if err != nil {
		return err
	}

	httpServer := &http.Server{
		Addr:    viper.GetString("backend.port"),
		Handler: r,
	}

	log.Info("app is successfully running on port " + viper.GetString("backend.port"))

	go func() {
		err = httpServer.ListenAndServe()
		if err != nil {
			done <- struct{}{}
			log.Fatal(err)
		}
	}()

	ready <- struct{}{}

	return nil
}

func extractAuthorizationParams(req *http.Request) *acsmw.InputParams {
	return &acsmw.InputParams{
		Login:       sysuser.GetLoginFromContext(req.Context()),
		Route:       req.RequestURI,
		Namespace:   warmup.Namespace,
		System:      warmup.SystemName,
		RouteMethod: req.Method,
	}
}

func (s *Server) Start(r chi.Router) error {

	pg, err := storage.NewPGConnection(viper.GetString("database.postgres"))
	if err != nil {
		return err
	}

	pgx, err := storage.WrapPgConnWithSqlx(pg)
	if err != nil {
		return err
	}

	err = storage.RunMigrations(pg, "./migrations")
	if err != nil {
		return err
	}

	s.repository.user = user.NewRepository(pgx)
	s.repository.tag = tag.NewRepository(pgx)
	s.repository.route = route.NewRepository(pgx)
	s.repository.group = group.NewRepository(pgx)
	s.repository.grouproute = grouproute.NewRepository(pgx)
	s.repository.usergroup = usergroup.NewRepository(pgx)
	s.repository.userroute = userroute.NewRepository(pgx)
	s.repository.namespace = namespace.NewRepository(pgx)
	s.repository.system = system.NewRepository(pgx)
	s.repository.sysuser = sysuser.NewSecurityGuard(sysuser.NewRepository(pgx))
	s.repository.acs = acs.NewRepository(pgx)

	sessionService := session.NewSessionService()

	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Mount(apiPathPrefix, r)

	host := "http://localhost" + viper.GetString("backend.port")
	swaggerUrl := "/swagger/source"
	r.Group(func(public chi.Router) {
		r.Group(func(private chi.Router) {
			private.Use(sysuser.Auth(sessionService, s.repository.sysuser))
			private.Use(acsmw.AccessCheck(host, extractAuthorizationParams))
			public.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(host+swaggerUrl)))
			public.Get(swaggerUrl, getSwaggerSource)

			sysuser.NewController(s.repository.sysuser, sessionService).Mount(private, public)
			acs.NewController(s.repository.user, s.repository.acs, s.repository.namespace, s.repository.system).Mount(public)

			namespace.NewController(s.repository.namespace).Mount(private)
			system.NewController(s.repository.system).Mount(private)
			tag.NewController(s.repository.tag).Mount(private)
			route.NewController(s.repository.route).Mount(private)
			group.NewController(s.repository.group).Mount(private)
			grouproute.NewController(s.repository.grouproute).Mount(private)
			user.NewController(s.repository.user, s.repository.group, s.repository.usergroup).Mount(private)
			usergroup.NewController(s.repository.usergroup).Mount(private)
			userroute.NewController(s.repository.userroute).Mount(private)
		})
	})

	return nil
}

func getSwaggerSource(writer http.ResponseWriter, _ *http.Request) {

	curDir, err := os.Getwd()
	if err != nil {
		return
	}

	file, err := os.Open(path.Join(curDir, "/generated/swagger.json"))
	if err != nil {
		return
	}
	defer func() {
		_ = file.Close()
	}()

	swaggerSpec, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	_, _ = writer.Write(swaggerSpec)
}

func buildRoutes(r chi.Router) []route.Route {
	rs := make([]route.Route, 0)
	for _, r := range r.Routes() {

		if apiPathPrefix+"/*" == r.Pattern {
			continue
		}

		rs = append(rs, route.Route{
			Route:       apiPathPrefix + r.Pattern,
			Method:      http.MethodPost,
			Description: "system",
		})
	}
	return rs
}
