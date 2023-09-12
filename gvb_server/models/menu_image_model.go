package models

// MenuBannerModel 自定义菜单和背景图的连接表 方便排序
type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id" `
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `json:"banner_id" `
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        uint        `json:"sort" gorm:"size:10"`
}
