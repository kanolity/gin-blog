package common

import (
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	db := global.DB
	if option.Debug {
		db = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	count = db.Select("id").Find(&list).RowsAffected
	offSet := (option.Page - 1) * option.Limit
	if offSet < 0 {
		offSet = 0
	}
	err = db.Limit(option.Limit).Offset(offSet).Find(&list).Error
	return list, count, err
}
