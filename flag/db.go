package flag

import (
	"go_blog/global"
	"go_blog/model"
)

func Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&model.UserModel{}, "CollectsModels", &model.UserCollectModel{})
	global.DB.SetupJoinTable(&model.MenuModel{}, "Banners", &model.MenuBannerModel{})

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.BannerModel{},
		&model.TagModel{},
		&model.MessageModel{},
		&model.AdvertModel{},
		&model.UserModel{},
		&model.CommentModel{},
		&model.ArticleModel{},
		&model.MenuModel{},
		&model.MenuBannerModel{},
		&model.FadeBackModel{},
		&model.LoginDataModel{},
	)

	if err != nil {
		global.Log.Error("[error]生成数据库表结构失败")
		return
	}
	global.Log.Info("[success]生成数据库表结构成功")
}
