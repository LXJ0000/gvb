package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/image_service"
)

var (
	// WhiteImageList 图片白名单
	WhiteImageList = []string{
		".jpg", ".png", ".jpeg", ".ico", ".tiff", ".gif", ".svg", ".webp",
	}
)

// ImagesUploadView 上传多个图片 返回图片的URL
func (ImagesApi) ImagesUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	files, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("文件不存在", c)
		return
	}

	// 不存在就创建
	var resList []image_service.FileUploadResponse

	for _, file := range files {
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		if !global.Config.QiNiu.Enable {
			//本地存储
			_ = c.SaveUploadedFile(file, serviceRes.FilePath)
		}
		resList = append(resList, serviceRes)

	}
	res.OKWithData(resList, c)
}
