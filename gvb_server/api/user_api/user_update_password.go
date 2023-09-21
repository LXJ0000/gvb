package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"`
	Pwd    string `json:"pwd" binding:"required" msg:"请输入新密码"`
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var user models.UserModel
	if err := global.DB.First(&user, userID).Error; err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	if !utils.CheckPassword(cr.OldPwd, user.Password) {
		res.FailWithMessage("密码错误", c)
		return
	}
	hashPwd := utils.SHA1(cr.Pwd)
	if err := global.DB.Where(&user).Update("password", hashPwd).Error; err != nil {
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OKWithC(c)
}
