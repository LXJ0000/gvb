package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	id := c.Param("id")
	//先把banner清空
	var menuModel models.MenuModel
	if err := global.DB.Preload("Banners").Take(&menuModel, id).Error; err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	global.DB.Where("menu_id = ?", id).Delete(&models.MenuBannerModel{})

	//如果选择了banner 则添加
	if len(cr.ImageSortList) > 0 {
		//添加banner

		var menuBannerList []models.MenuBannerModel

		for _, sort := range cr.ImageSortList {
			//判断image_id是否真正有这张图片 todo
			menuBannerList = append(menuBannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}

		if err := global.DB.Create(&menuBannerList).Error; err != nil {
			global.Log.Error(err)
			res.FailWithMessage("菜单图片关联失败", c)
			return
		}
	}
	//	普通更新
	maps := structs.Map(&cr)
	if err := global.DB.Model(&menuModel).Updates(maps).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单修改失败", c)

		return
	}
	res.OKWithMessage("菜单修改成功", c)
}
