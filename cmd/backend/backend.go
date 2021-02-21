package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	acsmw "github.com/hardstylez72/bzdacs-go"
	"github.com/hardstylez72/bzdacs/pkg/acs"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/grouproute"
	"github.com/hardstylez72/bzdacs/pkg/infra/logger"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/namespace"
	"github.com/hardstylez72/bzdacs/pkg/route"
	"github.com/hardstylez72/bzdacs/pkg/system"
	"github.com/hardstylez72/bzdacs/pkg/tag"
	"github.com/hardstylez72/bzdacs/pkg/user"
	"github.com/hardstylez72/bzdacs/pkg/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/userroute"
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
	}
}

func (s *Server) startBackendServer(log *zap.SugaredLogger, done chan struct{}) error {
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
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    viper.GetString("backend.port"),
		Handler: r,
	}

	log.Info("app is successfully running on port " + viper.GetString("backend.port"))

	err = httpServer.ListenAndServe()
	if err != nil {
		done <- struct{}{}
		log.Fatal(err)
	}
	return nil
}

func extractLogin(req *http.Request) string {
	return user.GetLoginFromContext(req.Context())
}

func extractRouteAndMethod(req *http.Request) (route, method string) {
	return req.URL.Path, req.Method
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

	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Mount(apiPathPrefix, r)

	host := "http://localhost" + viper.GetString("backend.port")
	r.Group(func(public chi.Router) {
		r.Group(func(private chi.Router) {
			swaggerUrl := "/swagger/source"
			r.Get("/swagger/*", httpSwagger.Handler(
				httpSwagger.URL(host+swaggerUrl), //The url pointing to API definition"
			))
			r.Get(swaggerUrl, getSwaggerSource)
			private.Use(user.Auth())
			private.Use(acsmw.AccessCheck(acsmw.NewService(host), extractLogin, extractRouteAndMethod))
			namespace.NewController(s.repository.namespace).Mount(private)
			system.NewController(s.repository.system).Mount(private)
			acs.NewController(s.repository.userroute, s.repository.user, s.repository.usergroup).Mount(public)
			tag.NewController(s.repository.tag).Mount(private)
			route.NewController(s.repository.route).Mount(private)
			group.NewController(s.repository.group).Mount(private)
			grouproute.NewController(s.repository.grouproute).Mount(private)
			user.NewController(s.repository.user, s.repository.group, s.repository.usergroup).Mount(private, public)
			usergroup.NewController(s.repository.usergroup).Mount(private)
			userroute.NewController(s.repository.userroute).Mount(private)
		})
	})

	ctx := context.Background()

	force := false
	err = s.initialize(ctx, force, r)
	if err != nil {
		return err
	}

	return nil
}

func getSwaggerSource(writer http.ResponseWriter, request *http.Request) {

	curDir, err := os.Getwd()
	if err != nil {
		return
	}

	file, err := os.Open(path.Join(curDir, "/docs/swagger.yaml"))
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
