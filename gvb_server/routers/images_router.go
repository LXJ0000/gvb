package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func ImagesRouter(router *gin.RouterGroup) {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("/images/", imagesApi.ImagesUploadView)
}
