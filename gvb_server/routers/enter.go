package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	apiRouterGroup := router.Group("api")

	//系统配置API
	SettingsRouter(apiRouterGroup)
	return router
}
