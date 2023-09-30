package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (UserApi) UserDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var userList models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 || len(cr.IDList) == 0 {
		//	不存在
		res.FailWithMessage("用户不存在", c)
		return
	}

	//事务
	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		// todo 用户消息表 用户收藏文章表 用户评论表 用户发布文章
		if err := tx.Delete(&userList).Error; err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	}); err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除用户失败", c)
		return
	}

	res.OKWithMessage(fmt.Sprintf("共删除 %d 个用户", count), c)
}
