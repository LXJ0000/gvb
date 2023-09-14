package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func SettingsRouter(router *gin.RouterGroup) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings/:name/", settingsApi.SettingsInfoView)
	router.POST("/settings/:name/", settingsApi.SettingsUpdateView)
}
