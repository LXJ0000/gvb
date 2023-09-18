package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	//先查菜单 取出菜单ID
	var menuList []models.MenuModel
	var menuIDList []uint
	//从小到大排序
	global.DB.Order("sort").Find(&menuList).Select("id").Scan(&menuIDList)

	//再查连接表
	var menuBanners []models.MenuBannerModel
	//从小到大排序
	global.DB.Preload("BannerModel").Order("sort").Find(&menuBanners, "menu_id in ?", menuIDList)

	var menuResponse []MenuResponse
	for _, menu := range menuList {
		//
		var banners = make([]Banner, 0)
		for _, banner := range menuBanners {
			if menu.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menuResponse = append(menuResponse, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}
	res.OKWithData(menuResponse, c)
}
