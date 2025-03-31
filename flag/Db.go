package flag

import (
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
)

// MakeMigration 表结构迁移
func MakeMigration() {
	var err error
	global.DB.SetupJoinTable(&models.User{}, "CollectsModels", &models.UserCollects{})
	global.DB.SetupJoinTable(&models.Menu{}, "Banners", &models.MenuBanner{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.Ad{},
			&models.Article{},
			&models.Banner{},
			&models.Comment{},
			&models.Feedback{},
			&models.LoginData{},
			&models.Menu{},
			&models.Message{},
			&models.Tag{},
			&models.User{},
			&models.ArticleTag{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功")
}
