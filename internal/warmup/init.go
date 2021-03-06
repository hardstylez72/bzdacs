package warmup

import (
	"context"
	"github.com/hardstylez72/bzdacs/config"
	sysuser "github.com/hardstylez72/bzdacs/internal/user"
	"github.com/hardstylez72/bzdacs/pkg/infra/storage"
	"github.com/hardstylez72/bzdacs/pkg/route"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var (
	SystemName                 string
	Namespace                  string
	AdminGroupName             string
	AdminGroupDescription      string
	GuestGroupName             string
	GuestGroupDescription      string
	HasGuest                   bool
	AdminLogin                 string
	AdminPassword              string
	SessionExpirationInSeconds int
)

type Config struct {
	Host     string
	BasePath string
	Login    string
	Password string
	Timeout  time.Duration
}

func Init(ctx context.Context, routes []route.Route) error {

	SystemName = viper.GetString("app.system")
	Namespace = viper.GetString("app.namespace")
	AdminGroupName = viper.GetString("app.adminGroupName")
	AdminGroupDescription = viper.GetString("app.adminGroupDescription")
	GuestGroupName = viper.GetString("app.guestGroupName")
	GuestGroupDescription = viper.GetString("app.guestGroupDescription")
	HasGuest = viper.GetBool("app.hasGuest")
	AdminLogin = viper.GetString("user.login")
	AdminPassword = viper.GetString("user.password")
	SessionExpirationInSeconds = viper.GetInt("user.sessionExpirationInSeconds")

	pg, err := storage.NewPGConnection(config.GetPostgresConn())
	if err != nil {
		return err
	}

	pgx, err := storage.WrapPgConnWithSqlx(pg)
	if err != nil {
		return err
	}
	defer func() { _ = pgx.Close() }()

	sysUserRepository := sysuser.NewSecurityGuard(sysuser.NewRepository(pgx))

	err = registerUser(ctx, sysUserRepository, AdminLogin, AdminPassword)
	if err != nil {
		return err
	}

	s, err := resolveSystem(ctx, pgx, SystemName)
	if err != nil {
		return err
	}

	ns, err := resolveNamespace(ctx, pgx, Namespace, s.Id)
	if err != nil {
		return err
	}

	u, err := resolveUser(ctx, pgx, AdminLogin, ns.Id)
	if err != nil {
		return err
	}

	rs, err := resolveRoutes(ctx, pgx, routes, ns.Id)
	if err != nil {
		return err
	}

	g, err := resolveGroup(ctx, pgx, AdminGroupName, AdminGroupDescription, ns.Id)
	if err != nil {
		return err
	}

	err = resolveGroupAndRoutes(ctx, pgx, g, rs)
	if err != nil {
		return err
	}

	err = resolveUserAndGroup(ctx, pgx, g, u)
	if err != nil {
		return err
	}

	if HasGuest {
		guestRoutes, err := resolveRoutes(ctx, pgx, filterGuestRoutes(routes), ns.Id)
		if err != nil {
			return err
		}
		guestGroup, err := resolveGroup(ctx, pgx, GuestGroupName, GuestGroupDescription, ns.Id)
		if err != nil {
			return err
		}
		err = resolveGroupAndRoutes(ctx, pgx, guestGroup, guestRoutes)
		if err != nil {
			return err
		}
	}

	return nil
}

func filterGuestRoutes(in []route.Route) []route.Route {
	out := make([]route.Route, 0)

	for _, r := range in {
		if strings.Contains(r.Route, "/update") ||
			strings.Contains(r.Route, "/delete") ||
			strings.Contains(r.Route, "/create") {
			continue
		}
		out = append(out, r)
	}
	return out
}
