package models

import "gorm.io/gorm"

type TagModel struct {
	gorm.Model
	Title   string         `json:"title" gorm:"size:16"`
	Article []ArticleModel `json:"-" gorm:"many2many:article_tag_models"`
}
