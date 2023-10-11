package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `json:"name" gorm:"size:32"`
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片类型 本地还是七牛云
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//还需删除本地存储
		if err := os.Remove(b.Path[1:]); err != nil {
			global.Log.Error(err.Error())
			return err
		}
	}
	return nil
}
