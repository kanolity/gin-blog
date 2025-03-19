package models

import "time"

// UserCollects 自定义表，记录用户收藏的时间
type UserCollects struct {
	UserID       uint    `gorm:"primaryKey"`
	UserModel    User    `gorm:"foreignkey:UserID"`
	ArticleID    uint    `gorm:"primaryKey"`
	ArticleModel Article `gorm:"foreignkey:ArticleID"`
	CreateTime   time.Time
}
