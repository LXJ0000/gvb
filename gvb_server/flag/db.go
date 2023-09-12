package flags

import (
	"gvb_server/global"
	"gvb_server/models"
)

func MakeMigrations() {
	var err error
	_ = global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	_ = global.DB.SetupJoinTable(&models.UserModel{}, "Banners", &models.MenuBannerModel{})

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.AdverModel{},
			&models.ArticleModel{},
			&models.BannerModel{},
			&models.CommentModel{},
			&models.FadeBackModel{},
			&models.MenuBannerModel{},
			&models.MenuModel{},
			&models.MessageModel{},
			&models.TagModel{},
			&models.UserModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功！")
}
