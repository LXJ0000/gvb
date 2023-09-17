package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImagesResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// ImagesNameListView 图片名称列表
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/images_names/ [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]ImagesResponse}
func (ImagesApi) ImagesNameListView(c *gin.Context) {

	var imageList []ImagesResponse

	global.DB.Model(&models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)

	res.OKWithData(imageList, c)

}
