package models

import "gorm.io/gorm"

// AdvertModel 广告表
type AdvertModel struct {
	gorm.Model
	Title  string `json:"title" gorm:"size:32"` //
	Href   string `json:"href"`                 // 链接
	Images string `json:"images"`               //
	IsShow bool   `json:"is_show"`              // 是否展示
}
