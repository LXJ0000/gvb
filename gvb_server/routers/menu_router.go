package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func MenuRouter(router *gin.RouterGroup) {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("/menu/", menuApi.MenuCreateView)
}
