package adverts_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法"`
	Images string `json:"images" binding:"required,url" msg:"图片地址非法"`
	IsShow bool   `json:"is_show" binding:"required" msg:"请选择是否展示"`
}

func (AdvertsApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest

	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//重复判断
	if err := global.DB.Select("id").Where("title = ?", cr.Title).First(&models.AdvertModel{}); err != nil {
		res.FailWithMessage("该广告已存在", c)
		return
	}

	if err := global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage("广告添加失败", c)
	}
	res.OKWithMessage("广告添加成功", c)
}
