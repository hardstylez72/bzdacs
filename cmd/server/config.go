package main

import (
	"github.com/spf13/viper"
)

func Load(filePath string) error {

	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}

	return nil
}
