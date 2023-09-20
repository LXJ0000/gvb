package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/desens"
)

type UserResponse struct {
	models.UserModel
}

func (UserApi) UserListView(c *gin.Context) {
	role, ok := c.Get("role")
	if !ok {
		res.FailWithMessage("", c)
		return
	}

	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	for _, user := range list {
		if role != int(ctype.PermissionAdmin) {
			// not Admin
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, user)
	}

	res.OKWithList(users, count, c)
}
