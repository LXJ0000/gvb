package models

// AdvertModel 广告表
type AdvertModel struct {
	MODEL
	Title  string `json:"title" gorm:"size:32"` //
	Href   string `json:"href"`                 // 链接
	Images string `json:"images"`               //
	IsShow bool   `json:"is_show"`              // 是否展示
}
