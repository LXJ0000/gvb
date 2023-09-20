package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type UserRoleUpdateRequest struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	UserID   uint       `json:"user_id" binding:"required" msg:"请选择正确的用户ID"`
	NickName string     `json:"nick_name"`
}

// UserUpdateRoleView 修改用户权限
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRoleUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	if err := global.DB.First(&user, cr.UserID).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage("用户不存在", c)
		return
	}
	if err := global.DB.Model(&user).Updates(map[string]interface{}{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage("更新失败", c)
		return
	}
	res.OKWithC(c)
}
