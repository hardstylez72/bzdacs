package internal

import (
	"context"
	libc "github.com/go-openapi/runtime/client"
	"github.com/hardstylez72/bzdacs/generated/client"
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

func Init(ctx context.Context, routes *[]route.Route) error {

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

	c := GetClient()

	err := registerUser(ctx, c, AdminLogin, AdminPassword)
	if err != nil {
		return err
	}

	s, err := resolveSystem(ctx, c, SystemName)
	if err != nil {
		return err
	}

	ns, err := resolveNamespace(ctx, c, Namespace, s.Id)
	if err != nil {
		return err
	}

	u, err := resolveUser(ctx, c, AdminLogin, ns.Id)
	if err != nil {
		return err
	}

	rs, err := resolveRoutes(ctx, c, *routes, ns.Id)
	if err != nil {
		return err
	}

	g, err := resolveGroup(ctx, c, AdminGroupName, AdminGroupDescription, ns.Id)
	if err != nil {
		return err
	}

	err = resolveGroupAndRoutes(ctx, c, g, rs)
	if err != nil {
		return err
	}

	err = resolveUserAndGroup(ctx, c, g, u)
	if err != nil {
		return err
	}

	if HasGuest {
		guestRoutes, err := resolveRoutes(ctx, c, filterGuestRoutes(*routes), ns.Id)
		if err != nil {
			return err
		}
		guestGroup, err := resolveGroup(ctx, c, GuestGroupName, GuestGroupDescription, ns.Id)
		if err != nil {
			return err
		}
		err = resolveGroupAndRoutes(ctx, c, guestGroup, guestRoutes)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetConfig() *Config {
	return &Config{
		Host:     "localhost:4000",
		BasePath: "/api",
		Login:    AdminLogin,
		Password: AdminPassword,
		Timeout:  time.Second * 10,
	}
}

func GetClient() *client.BZDACS {
	transport := libc.New(GetConfig().Host, GetConfig().BasePath, nil)
	transport.DefaultAuthentication = libc.BasicAuth(GetConfig().Login, GetConfig().Password)
	return client.New(transport, nil)
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
