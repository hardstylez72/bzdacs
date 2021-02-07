package config

import (
	"github.com/spf13/viper"
	"os"
)

func Load(filePath string) error {

	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}

	if viper.Get("env") == "prod" {
		login := os.Getenv("BZDACS_LOGIN")
		if login != "" {
			viper.Set("user.login", login)
		}
		password := os.Getenv("BZDACS_PASSWORD")
		if password != "" {
			viper.Set("user.password", password)
		}

		postgres := os.Getenv("$BZDACS_POSTGRES")
		if postgres != "" {
			viper.Set("database.postgres", postgres)
		}
	}

	return nil
}
