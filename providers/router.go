package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"xd_working_trial/cfg"
	"xd_working_trial/ginServer"
	"xd_working_trial/handlers"
)

// setupRouter setup router.
func setupRouter(hs *handlers.Handlers) ginServer.GinRoutingFn {
	return func(router *gin.Engine) {
		baseRoute := router.Group(viper.GetString(cfg.ConfigKeyContextPath))
		baseRoute.GET("/health-check", hs.HealthCheck.HealthCheck())
		info := baseRoute.Group("/info")
		{
			info.GET("/os", hs.AppHandler.GetOSInfo())
			// info.GET("/user", hs.AppHandler.GetUserInfo())
			// info.GET("/metric", hs.AppHandler.GetMetricInfo())
		}
	}
}
