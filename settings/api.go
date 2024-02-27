package settings

import "time"

func GetWebServerAddress() string {
	return config.GetString("address")+":"+config.GetString("port")
}

func GetMode() string {
	return config.GetString("mode")
}

func ShutdownTimeout() time.Duration {
	return 1
}