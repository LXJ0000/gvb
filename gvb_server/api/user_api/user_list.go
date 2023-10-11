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
	RoleID int `json:"role_id"`
}

type UserListRequest struct {
	models.PageInfo
	Role int `json:"role" form:"role"`
}

func (UserApi) UserListView(c *gin.Context) {
	role, ok := c.Get("role")
	if !ok {
		res.FailWithMessage("", c)
		return
	}

	var cr UserListRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []UserResponse
	list, count, _ := common.ComList(models.UserModel{Role: ctype.Role(cr.Role)}, common.Option{
		PageInfo: cr.PageInfo,
		Likes:    []string{"nick_name"},
	})

	for _, user := range list {
		if role != int(ctype.PermissionAdmin) {
			// not Admin
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
	}

	res.OKWithList(users, count, c)
}
