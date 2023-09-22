package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/service"
)

// UserLogoutView 注销
func (UserApi) UserLogoutView(c *gin.Context) {

	token := c.Request.Header.Get("token")
	expRaw, _ := c.Get("ExpiresAt")

	service.ServiceApp.USerService.Logout(expRaw, token)

	res.OKWithC(c)
}
