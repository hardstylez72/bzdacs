package config

import "github.com/spf13/viper"

type internal struct {
	AdminLogin                 string
	AdminPassword              string
	SessionExpirationInSeconds int
	System                     string
	Namespace                  string
	AdminGroup                 Group
	GuestGroup                 Group
	HasGuest                   bool
}

type Group struct {
	Name        string
	Description string
}

func GetInternal() internal {
	return internal{
		AdminLogin:                 viper.GetString("internal.login"),
		AdminPassword:              viper.GetString("internal.password"),
		SessionExpirationInSeconds: viper.GetInt("internal.sessionExpirationInSeconds"),
		System:                     viper.GetString("internal.system"),
		Namespace:                  viper.GetString("internal.namespace"),
		AdminGroup: Group{
			Name:        viper.GetString("internal.adminGroupName"),
			Description: viper.GetString("internal.adminGroupDescription"),
		},
		GuestGroup: Group{
			Name:        viper.GetString("internal.guestGroupName"),
			Description: viper.GetString("internal.guestGroupDescription"),
		},
		HasGuest: viper.GetBool("internal.hasGuest"),
	}
}
