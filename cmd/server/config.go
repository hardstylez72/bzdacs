package main

import (
	"github.com/hardstylez72/bblog/internal/auth"
	"github.com/hardstylez72/bblog/internal/objectstorage"
	"github.com/spf13/viper"
)

const (
	cfgAllData = ""
)

type Config struct {
	Port      string
	Env       string
	Host      string
	Oauth     auth.Oauth
	Databases Databases
	ObjectStorage
	SessionCookie auth.SessionCookieConfig
}

type ObjectStorage struct {
	Minio objectstorage.Config
}

type Databases struct {
	Postgres string
}

func Load(filePath string) error {

	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}

	return nil
}
