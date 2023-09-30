package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/utils"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string     `json:"password" binding:"required" msg:"请输入密码"`
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var user models.UserModel
	if err := global.DB.Where("user_name=?", cr.UserName).First(&user); err == nil {
		res.FailWithMessage("用户名已存在", c)
		return
	}

	cr.Password = utils.SHA1(cr.Password)
	if err := global.DB.Create(&models.UserModel{
		NickName: cr.NickName,
		UserName: cr.UserName,
		Password: cr.Password,
		Role:     cr.Role,
		Avatar:   "/static/avatar/2.jpg",
		IP:       c.ClientIP(),
	}).Error; err != nil {
		res.FailWithMessage("用户创建失败", c)
	}
	res.OKWithMessage(fmt.Sprintf("用户%s创建成功", cr.UserName), c)
}
