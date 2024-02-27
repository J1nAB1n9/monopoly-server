package httpserver

import (
	"github.com/gin-gonic/gin"
	"monopoly-server/settings"
)

func routerEngine() *gin.Engine {
	gin.SetMode(settings.GetMode())

	r := gin.New()
	r.Use(gin.Recovery())



	return r
}