package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// ImagesListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/images/ [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
func (ImagesApi) ImagesListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Likes:    []string{"name"},
	})

	res.OKWithList(list, count, c)

}
