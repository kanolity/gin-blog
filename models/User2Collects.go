package models

import "time"

// User2Collects 自定义表，记录用户收藏的时间
type User2Collects struct {
	UserID       uint    `gorm:"primaryKey"`
	UserModel    User    `gorm:"foreignkey:UserID"`
	ArticleID    uint    `gorm:"primaryKey"`
	ArticleModel Article `gorm:"foreignkey:ArticleID"`
	CreateTime   time.Time
}
