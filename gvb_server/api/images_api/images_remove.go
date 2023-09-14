package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (ImagesApi) ImagesRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imagesList []models.BannerModel
	count := global.DB.Find(&imagesList, cr.IDList).RowsAffected
	if count == 0 {
		//	不存在
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Delete(&imagesList)
	res.OKWithMessage(fmt.Sprintf("共删除%d张图片", count), c)
}
