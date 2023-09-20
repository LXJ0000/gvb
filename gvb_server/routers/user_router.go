package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func UserRouter(router *gin.RouterGroup) {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/email_login/", userApi.EmailLoginView)
	router.GET("/user/", middleware.JwtAuth(), userApi.UserListView)
	router.POST("/user_role/", middleware.JwtAdminAuth(), userApi.UserUpdateRoleView)

}
