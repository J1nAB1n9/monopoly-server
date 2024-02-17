package settings

import "github.com/spf13/viper"

func NewConfig() *viper.Viper {
	viper.SetConfigName("config")
	viper.SetConfigFile(".")
	viper.SetConfigType("toml")

	return viper.GetViper()
}