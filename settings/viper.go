package settings

import (
	"github.com/spf13/viper"
	"monopoly-server/logger"
)

var config *viper.Viper

func GetConfig() *viper.Viper {
	return config
}
func NewConfig()  {
	config = viper.New()
	config.SetConfigType("toml")
	config.AddConfigPath("./config")

	// default
	config.SetDefault("mode","debug")
	config.SetDefault("address","127.0.0.1")
	config.SetDefault("port" , "8080")

	// init viper
	if err := config.ReadInConfig();err != nil {
		logger.Fatal("loading config.toml error :%v",err)
	}
}