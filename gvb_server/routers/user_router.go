package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func UserRouter(router *gin.RouterGroup) {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/email_login/", userApi.EmailLoginView)

}
