package models

import (
	"gorm.io/gorm"
	"gvb_server/models/ctype"
)

// MenuModel 菜单的路径可以是 /path 也可以是路由别名
type MenuModel struct {
	gorm.Model
	Title        string        `gorm:"size:32" json:"title"`                                                                         //标题
	Path         string        `gorm:"size:32" json:"path"`                                                                          //路径
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                        //
	Abstract     ctype.Array   `gorm:"" json:"abstract"`                                                                             //简介
	AbstractTime int           `json:"abstract_time"`                                                                                //简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:BannerID;JoinReferences:BannerID " json:"banners"` //菜单的图片列表
	BannerTime   int           ` json:"banner_time"`                                                                                 //菜单图片切换时间 0为不切换
	Sort         int           `gorm:"size:10" json:"sort"`                                                                          //菜单的顺序
}
