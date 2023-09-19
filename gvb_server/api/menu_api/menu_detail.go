package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	var menu models.MenuModel
	id := c.Param("id")
	//从小到大排序
	//Find找不到返回空 不会返回错误 而First则会
	if err := global.DB.First(&menu, id).Error; err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	//再查连接表
	var menuBanners []models.MenuBannerModel
	//从小到大排序
	global.DB.Preload("BannerModel").Order("sort").Find(&menuBanners, "menu_id = ?", id)

	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menu,
		Banners:   banners,
	}
	res.OKWithData(menuResponse, c)
}
