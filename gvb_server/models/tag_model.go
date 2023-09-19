package models

type TagModel struct {
	MODEL
	Title   string         `json:"title" gorm:"size:16"`
	Article []ArticleModel `json:"-" gorm:"many2many:article_tag_models"`
}
