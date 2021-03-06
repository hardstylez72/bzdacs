package config

import (
	"github.com/spf13/viper"
	"strings"
)

type backend struct {
	Host string
	Port string
}

func GetBackend() backend {
	return backend{
		Host: "http://localhost" + viper.GetString("backend.port"),
		Port: viper.GetString("backend.port"),
	}
}

func GetProxy() backend {
	return backend{
		Host: "http://localhost" + viper.GetString("proxy.port"),
		Port: viper.GetString("proxy.port"),
	}
}

func GetPostgresConn() string {
	return viper.GetString("database.postgres")
}

type staticServer struct {
	Port string
	Host string
	Path string
}

func GetStaticServer() staticServer {
	return staticServer{
		Port: viper.GetString("ui.port"),
		Path: viper.GetString("ui.path"),
		Host: "http://localhost" + viper.GetString("ui.port"),
	}
}

type app struct {
	Host   string
	Domain string
}

func GetApp() app {
	domain := strings.ReplaceAll(viper.GetString("host"), "http://", "")
	domain = strings.ReplaceAll(domain, "https://", "")

	return app{
		Host:   viper.GetString("host"),
		Domain: domain,
	}
}
