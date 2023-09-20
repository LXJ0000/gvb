package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//	判断用户是否存在
	var user models.UserModel
	if err := global.DB.First(&user, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error; err != nil {
		res.FailWithMessage("用户名或密码错误", c)
		global.Log.Warnln("用户名or邮箱 不存在")
		return
	}

	//密码校验
	if hashPassword := utils.SHA1(cr.Password); hashPassword != user.Password {
		res.FailWithMessage("用户名或密码错误", c)
		global.Log.Warnln("密码不正常")
		return
	}

	//	token生成
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   user.ID,
		NickName: user.NickName,
		Role:     int(user.Role),
		Avatar:   user.Avatar,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", c)
		return
	}
	res.OKWithData(token, c)
}
