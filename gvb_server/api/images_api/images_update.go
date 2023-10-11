package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImagesUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

// ImagesUpdateView 图片更新
// @Tags 图片管理
// @Summary 图片更新
// @Description 图片更新
// @Param data body ImagesUpdateRequest   true  "表示多个参数"
// @Router /api/images/ [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImagesUpdateView(c *gin.Context) {
	var cr ImagesUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var imageModel models.BannerModel
	if err := global.DB.Where("id=?", cr.ID).Find(&imageModel).Error; err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}
	if err := global.DB.Model(&imageModel).Update("name", cr.Name).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OKWithMessage("图片名称修改成功", c)

}
