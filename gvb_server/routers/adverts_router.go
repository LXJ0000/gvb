package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func AdvertRouter(router *gin.RouterGroup) {
	advertApi := api.ApiGroupApp.AdvertsApi
	router.POST("adverts", advertApi.AdvertCreateView)
}
