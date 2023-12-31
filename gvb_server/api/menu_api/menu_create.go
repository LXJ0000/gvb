package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请输入菜单中文名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请输入菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`
	BannerTime    int         `json:"banner_time" structs:"banner_time"`
	Sort          int         `json:"sort" binding:"required" msg:"请选择菜单序号" structs:"sort"`
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"` // 具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	//判断重复
	var count int64
	if global.DB.Model(&models.MenuModel{}).Where("title = ? or path = ?", cr.Title, cr.Path).Count(&count); count > 0 {
		res.FailWithMessage("此菜单已存在", c)
		return
	}
	//	menu入库
	if err := global.DB.Create(&menuModel).Error; err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("菜单添加失败", c)
		return
	}

	if len(cr.ImageSortList) == 0 {
		res.FailWithMessage("菜单添加成功", c)
		return
	}

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
	res.OKWithMessage("菜单添加成功", c)
}
