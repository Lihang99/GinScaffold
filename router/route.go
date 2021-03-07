package router

import (
	"GinScaffold/logger"
	"GinScaffold/middlewares"
	"GinScaffold/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "404"})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	if settings.Conf.JwtConfig.Enable {
		r.Use(middlewares.JWTAuthMiddleware())
	}
	return r
}
