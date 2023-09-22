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
	router.DELETE("/user/", middleware.JwtAdminAuth(), userApi.UserDeleteView)

	router.POST("/user_role/", middleware.JwtAdminAuth(), userApi.UserUpdateRoleView)
	router.POST("/user_password/", middleware.JwtAuth(), userApi.UserUpdatePassword)
	router.POST("/logout/", middleware.JwtAuth(), userApi.UserLogoutView)

}
