package internal

import (
	"context"
	libc "github.com/go-openapi/runtime/client"
	"github.com/hardstylez72/bzdacs/client"
	"github.com/hardstylez72/bzdacs/pkg/route"
	"time"
)

const (
	SystemName       = "BZDACS"
	Namespace        = "internal"
	GroupName        = "system"
	GroupDescription = "BZDACS system group"
	User             = "admin"
)

type Config struct {
	Host     string
	BasePath string
	Login    string
	Password string
	Timeout  time.Duration
}

func Init(ctx context.Context, routes *[]route.Route) error {

	c := GetClient()
	s, err := resolveSystem(ctx, c, SystemName)
	if err != nil {
		return err
	}

	ns, err := resolveNamespace(ctx, c, Namespace, s.Id)
	if err != nil {
		return err
	}

	_, err = resolveUser(ctx, c, User, ns.Id)
	if err != nil {
		return err
	}

	rs, err := resolveRoutes(ctx, c, *routes, ns.Id)
	if err != nil {
		return err
	}

	g, err := resolveGroup(ctx, c, GroupName, GroupDescription, ns.Id)
	if err != nil {
		return err
	}

	err = resolveGroupAndRoutes(ctx, c, g, rs)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return &Config{
		Host:     "localhost:4000",
		BasePath: "/api",
		Login:    "hs",
		Password: "1234",
		Timeout:  time.Second * 100,
	}
}

func GetClient() *client.BZDACS {
	transport := libc.New(GetConfig().Host, GetConfig().BasePath, nil)
	transport.DefaultAuthentication = libc.BasicAuth(GetConfig().Login, GetConfig().Password)
	return client.New(transport, nil)
}
