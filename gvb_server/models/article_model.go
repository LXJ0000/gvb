package models

import (
	"gvb_server/models/ctype"
)

type ArticleModel struct {
	MODEL
	Title        string         `json:"title" gorm:"size:32"`                          //标题
	Abstract     string         `json:"abstract" gorm:""`                              //简介
	Content      string         `json:"content" gorm:""`                               //内容
	LookCount    int            `json:"look_count" gorm:""`                            //浏览量
	CommentCount int            `json:"comment_count" gorm:""`                         //评论数
	DiggCount    int            `json:"digg_count" gorm:""`                            //点赞数
	CollectCount int            `json:"collect_count" gorm:""`                         //收藏数
	TagModel     []TagModel     `json:"tag_model" gorm:"many2many:article_tag_models"` //标签
	CommentModel []CommentModel `json:"-" gorm:"foreignKey:ArticleID"`                 //评论列表
	UserModel    UserModel      `json:"-" gorm:"foreignKey:UserID"`                    //作者
	UserID       uint           `json:"user_id" gorm:""`                               //用户ID
	Category     string         `json:"category" gorm:"size:20"`                       //分类
	Source       string         `json:"source" gorm:""`                                //来源
	Link         string         `json:"link" gorm:""`                                  //链接
	Banner       BannerModel    `json:"-" gorm:"foreignKey:BannerID"`                  //封面
	BannerID     uint           `json:"banner_id" `                                    //封面ID
	NickName     string         `json:"nick_name" gorm:"size:42"`                      //昵称
	BannerPath   string         `json:"banner_path" `                                  //封面地址
	Tags         ctype.Array    `json:"tags" gorm:"type:string;size:64"`               //标签
}
