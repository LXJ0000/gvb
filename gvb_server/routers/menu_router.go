package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func MenuRouter(router *gin.RouterGroup) {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("/menu/", menuApi.MenuCreateView)
	router.GET("/menu/", menuApi.MenuListView)
	router.GET("/menu_names/", menuApi.MenuNameListView)
	router.PUT("/menu/:id/", menuApi.MenuUpdateView)
	router.DELETE("/menu/", menuApi.MenuDeleteView)

}
