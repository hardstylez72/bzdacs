package main

import (
	"context"
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/hardstylez72/bzdacs/pkg/acs"
	"github.com/hardstylez72/bzdacs/pkg/group"
	"github.com/hardstylez72/bzdacs/pkg/grouproute"
	"github.com/hardstylez72/bzdacs/pkg/infra/logger"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/mwv"
	"github.com/hardstylez72/bzdacs/pkg/route"
	"github.com/hardstylez72/bzdacs/pkg/tag"
	"github.com/hardstylez72/bzdacs/pkg/user"
	"github.com/hardstylez72/bzdacs/pkg/usergroup"
	"github.com/hardstylez72/bzdacs/pkg/userroute"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

const (
	apiPathPrefix = "/api"
)

type Server struct {
	log    *zap.SugaredLogger
	router chi.Router
}

func main() {

	log, err := logger.New("")
	errCheck(err, "can't load config")
	defer log.Sync()

	err = NewServer(log).Run()
	errCheck(err, "can't run server")
}

func errCheck(err error, errorText string) {
	if err == nil {
		return
	}
	log.Fatal(errorText, ": ", err)
}

func NewServer(log *zap.SugaredLogger) *Server {
	return &Server{
		router: chi.NewRouter(),
		log:    log,
	}
}

func (s *Server) Run() error {

	configPath := flag.String("config", "/home/hs/go/src/github.com/hardstylez72/bzdacs/cmd/server/config.local.yaml", "path to config file")
	flag.Parse()

	err := Load(*configPath)
	if err != nil {
		return err
	}
	//err = tracer.New(viper.GetString("tracer.jaeger.collectorEndpoint"), viper.GetString("tracer.jaeger.serviceName"))
	//if err != nil {
	//	return err
	//}

	httpServer := &http.Server{
		Addr:    viper.GetString("port"),
		Handler: s.Handler(),
	}

	return httpServer.ListenAndServe()
}

func (s *Server) Handler() chi.Router {

	r := s.router
	c := cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	r.Use(c)
	r.Use(logger.Inject(s.log))

	err := Start(r)
	if err != nil {
		log.Fatal(err)
	}
	s.log.Info("app is successfully running")

	return r
}

func extractLogin(req *http.Request) string {
	return user.GetLoginFromContext(req.Context())
}

func extractRouteAndMethod(req *http.Request) (route, method string) {
	return req.URL.Path, req.Method
}

func Start(r chi.Router) error {
	pg, err := storage.NewPGConnection(viper.GetString("databases.postgres"))
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

	var (
		tagRepository        = tag.NewRepository(pgx)
		routeRepository      = route.NewRepository(pgx)
		groupRepository      = group.NewRepository(pgx)
		groupRouteRepository = grouproute.NewRepository(pgx)
		userRepository       = user.NewRepository(pgx)
		userGroupRepository  = usergroup.NewRepository(pgx)
		userRouteRepository  = userroute.NewRepository(pgx)
	)

	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Mount(apiPathPrefix, r)

	r.Group(func(public chi.Router) {
		r.Group(func(private chi.Router) {
			private.Use(user.Auth())
			private.Use(mwv.AccessCheck(mwv.NewService("http://localhost:3000"), extractLogin, extractRouteAndMethod))
			acs.NewController(userRouteRepository, userRepository, userGroupRepository).Mount(public)
			tag.NewController(tagRepository).Mount(private)
			route.NewController(routeRepository).Mount(private)
			group.NewController(groupRepository).Mount(private)
			grouproute.NewController(groupRouteRepository).Mount(private)
			user.NewController(userRepository, groupRepository, userGroupRepository).Mount(private, public)
			usergroup.NewController(userGroupRepository).Mount(private)
			userroute.NewController(userRouteRepository).Mount(private)
		})
	})

	ctx := context.Background()

	force := true

	u, err := resolveUser(ctx, userRepository, &Option{Force: force})
	if err != nil {
		return err
	}
	g, err := resolveGroup(ctx, groupRepository, &Option{Force: force})
	if err != nil {
		return err
	}
	err = resolveUserAndGroup(ctx, userGroupRepository, u.Id, g.Id, &Option{Force: force})
	if err != nil {
		return err
	}

	rs := buildRoutes(r)

	rs, err = resolveRoutes(ctx, routeRepository, rs, &Option{Force: force})
	if err != nil {
		return err
	}

	err = resolveGroupAndRoutes(ctx, groupRouteRepository, rs, g.Id, &Option{Force: force})
	if err != nil {
		return err
	}

	_, err = resolveGuestGroup(ctx, groupRepository, g.Id, &Option{Force: force})
	if err != nil {
		return err
	}

	return nil
}
