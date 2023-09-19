package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvb_server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	///swagger/index.html
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("/api/")

	//系统配置API todo
	SettingsRouter(apiRouterGroup)
	ImagesRouter(apiRouterGroup)
	AdvertRouter(apiRouterGroup)
	MenuRouter(apiRouterGroup)
	UserRouter(apiRouterGroup)
	return router
}
