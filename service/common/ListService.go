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
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按照时间倒序排列
	}
	query := db.Where(model)

	err = query.Model(&model).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	if option.Limit == 0 {
		option.Limit = 10
	}
	offSet := (option.Page - 1) * option.Limit
	if offSet < 0 {
		offSet = 0
	}
	err = query.Limit(option.Limit).Offset(offSet).Order(option.Sort).Find(&list).Error
	return list, count, err
}
