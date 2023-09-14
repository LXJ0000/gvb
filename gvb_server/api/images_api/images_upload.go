package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"path"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`
}

// ImagesUploadView 上传单张图片 返回图片的URL
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

	// 判断路径是否存在 todo 好像不需要 SaveUploadedFile会创建
	basePath := global.Config.Static.Path
	//
	// 不存在就创建
	var resList []FileUploadResponse

	for _, file := range files {
		//默认上传成功
		fileUploadResponse := FileUploadResponse{FileName: file.Filename, IsSuccess: true, Msg: "Success"}

		//上传路径
		filePath := path.Join(basePath, file.Filename)

		//判断大小
		if size := float64(file.Size) / float64(1024*1024); size > float64(global.Config.Static.Size) {
			//上传失败
			fileUploadResponse.IsSuccess = false
			fileUploadResponse.Msg = fmt.Sprintf("文件当前大小为：%.2fMb 超出限定大小：%dMb", size, global.Config.Static.Size)
		} else {
			err := c.SaveUploadedFile(file, filePath)
			if err != nil {
				global.Log.Info(err.Error())
				//上传失败
				fileUploadResponse.IsSuccess = false
				fileUploadResponse.Msg = err.Error()
			}
		}

		resList = append(resList, fileUploadResponse)
	}
	res.OKWithData(resList, c)
}
