package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var menuList models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		//	不存在
		res.FailWithMessage("菜单不存在", c)
		return
	}

	//事务
	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&menuList).Association("Banners").Clear(); err != nil {
			global.Log.Error(err)
			return err
		}
		if err := tx.Delete(&menuList).Error; err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	}); err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}

	res.OKWithMessage(fmt.Sprintf("共删除 %d 个菜单", count), c)
}
