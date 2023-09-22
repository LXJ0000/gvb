package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/utils/jwts"
	"net/http"
	"time"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("token")

		//用户不存在
		if tokenStr == "" {
			global.Log.Warnln("该用户不存在")
			c.JSON(http.StatusOK, res.Response{
				Code: 401,
				Msg:  "该用户不存在",
			})
			c.Abort() //阻止执行
			return
		}

		//验证token
		claim, err := jwts.ParseToken(tokenStr)
		if err != nil {
			global.Log.Warnln("token不正确")

			c.JSON(http.StatusOK, res.Response{
				Code: 403,
				Msg:  "token不正确",
			})
			c.Abort() //阻止执行
			return
		}

		//token超时
		if claim.ExpiresAt.Time.Before(time.Now()) {
			global.Log.Warnln("token过期")

			c.JSON(http.StatusOK, res.Response{
				Code: 402,
				Msg:  "token过期",
			})
			c.Abort() //阻止执行
			return
		}

		//判断是否在redis中 若是则注销了
		if service.ServiceApp.RedisService.CheckLogout(tokenStr) {
			global.Log.Warnln("token失效")

			c.JSON(http.StatusOK, res.Response{
				Code: 200, // todo modify
				Msg:  "token失效",
			})
			c.Abort() //阻止执行
			return
		}

		c.Set("role", claim.Role)
		c.Set("user_id", claim.UserID)
		c.Set("avatar", claim.Avatar)
		c.Set("nick_name", claim.NickName)
		c.Set("ExpiresAt", claim.ExpiresAt)
		c.Next()
	}
}

func JwtAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("token")

		//用户不存在
		if tokenStr == "" {
			global.Log.Warnln("该用户不存在")
			c.JSON(http.StatusOK, res.Response{
				Code: 401,
				Msg:  "该用户不存在",
			})
			c.Abort() //阻止执行
			return
		}

		//验证token
		claim, err := jwts.ParseToken(tokenStr)
		if err != nil {
			global.Log.Warnln("token不正确")

			c.JSON(http.StatusOK, res.Response{
				Code: 403,
				Msg:  "token不正确",
			})
			c.Abort() //阻止执行
			return
		}

		//token超时
		if claim.ExpiresAt.Time.Before(time.Now()) {
			global.Log.Warnln("token过期")

			c.JSON(http.StatusOK, res.Response{
				Code: 402,
				Msg:  "token过期",
			})
			c.Abort() //阻止执行
			return
		}

		//Admin ?
		if claim.Role != int(ctype.PermissionAdmin) {
			global.Log.Warnln("权限过低")

			c.JSON(http.StatusOK, res.Response{
				Code: 402,
				Msg:  "权限过低",
			})
			c.Abort() //阻止执行
			return
		}

		//判断是否在redis中 若是则注销了
		if service.ServiceApp.RedisService.CheckLogout(tokenStr) {
			global.Log.Warnln("token失效")

			c.JSON(http.StatusOK, res.Response{
				Code: 200, // todo modify
				Msg:  "token失效",
			})
			c.Abort() //阻止执行
			return
		}

		c.Set("role", claim.Role)
		c.Set("user_id", claim.UserID)
		c.Set("avatar", claim.Avatar)
		c.Set("nick_name", claim.NickName)

		c.Next()
	}
}
