package config

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
)

func Load(filePath string) error {

	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	if IsProd() {
		login := os.Getenv("BZDACS_LOGIN")
		if login != "" {
			viper.Set("user.login", login)
		}
		password := os.Getenv("BZDACS_PASSWORD")
		if password != "" {
			viper.Set("user.password", password)
		}
		sessionLifeTime := os.Getenv("BZDACS_SESSION_LIFETIME_IN_SECONDS")
		if sessionLifeTime != "" {
			inSeconds, err := strconv.Atoi(sessionLifeTime)
			if err != nil {
				return err
			}
			viper.Set("user.sessionExpirationInSeconds", inSeconds)
		}

		host := os.Getenv("BZDACS_HOST")
		if host != "" {
			viper.Set("host", host)
		}

		postgres := os.Getenv("BZDACS_POSTGRES")
		if postgres != "" {
			viper.Set("database.postgres", postgres)
		}
	}

	return nil
}

func IsProd() bool {
	return viper.GetString("env") == "prod"
}
