package tests

import (
	libc "github.com/go-openapi/runtime/client"
	"github.com/hardstylez72/bzdacs/client"
	"time"
)

type Config struct {
	Host     string
	BasePath string
	Login    string
	Password string
	Timeout  time.Duration
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
