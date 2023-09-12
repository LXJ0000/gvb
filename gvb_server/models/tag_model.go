package models

import "gorm.io/gorm"

type TagModel struct {
	gorm.Model
	Title    string         `json:"title" gorm:"size:16"`
	Ariticle []ArticleModel `json:"-" gorm:"many2many:article_tag"`
}
