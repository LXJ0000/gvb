package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"path"
	"path/filepath"
	"strings"
)

var (
	// WhiteImageList 图片白名单
	WhiteImageList = []string{
		".jpg", ".png", ".jpeg", ".ico", ".tiff", ".gif", ".svg", ".webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`
}

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

	// 判断路径是否存在 todo 好像不需要 SaveUploadedFile会创建
	basePath := global.Config.Static.Path
	//
	// 不存在就创建
	var resList []FileUploadResponse

	for _, file := range files {
		fileName := file.Filename
		//默认上传成功
		fileUploadResponse := FileUploadResponse{FileName: fileName, IsSuccess: true, Msg: "Success"}

		//上传路径
		filePath := path.Join(basePath, fileName)

		//文件名后缀
		suffix := strings.ToLower(filepath.Ext(fileName))
		if !utils.InList(suffix, WhiteImageList) {
			//上传失败
			fileUploadResponse.IsSuccess = false
			fileUploadResponse.Msg = "非法的文件格式"
		} else if size := float64(file.Size) / float64(1024*1024); size > float64(global.Config.Static.Size) {
			//上传失败
			fileUploadResponse.IsSuccess = false
			fileUploadResponse.Msg = fmt.Sprintf("文件当前大小为：%.2fMb 超出限定大小：%dMb", size, global.Config.Static.Size)
		} else {
			fileObj, _ := file.Open()
			if err != nil {
				global.Log.Error(err.Error())
			}
			byteData, err := io.ReadAll(fileObj)
			if err != nil {
				global.Log.Error(err.Error())
			}
			imageHash := utils.Md5(byteData)
			//查询数据库是否存在对应Hash
			var bannerModel models.BannerModel
			if err := global.DB.Where("hash = ?", imageHash).First(&bannerModel).Error; err == nil {
				//	找到了 不需要存入数据库
				fileUploadResponse.FileName = bannerModel.Path
				fileUploadResponse.IsSuccess = false
				fileUploadResponse.Msg = fmt.Sprintf("图片已存在")
			} else {
				if global.Config.QiNiu.Enable {
					//开启七牛云存储
					filePath, err = qiniu.UploadImage(byteData, fileName, "gvb")
					if err != nil {
						global.Log.Error(err.Error())
						fileUploadResponse.IsSuccess = false
						fileUploadResponse.Msg = err.Error()
					} else {
						//入库
						fileUploadResponse.Msg = "图片上传成功 ~七牛云"
						global.DB.Create(&models.BannerModel{
							Path:      filePath,
							Hash:      imageHash,
							Name:      fileName,
							ImageType: ctype.QiNiu,
						})
					}
				} else {
					if err := c.SaveUploadedFile(file, filePath); err != nil {
						global.Log.Info(err.Error())
						//上传失败
						fileUploadResponse.IsSuccess = false
						fileUploadResponse.Msg = err.Error()
					} else {
						//入库
						fileUploadResponse.Msg = "图片上传成功 ~本地"
						global.DB.Create(&models.BannerModel{
							Path:      filePath,
							Hash:      imageHash,
							Name:      fileName,
							ImageType: ctype.Local,
						})
					}
				}

			}
		}
		resList = append(resList, fileUploadResponse)
	}
	res.OKWithData(resList, c)
}
