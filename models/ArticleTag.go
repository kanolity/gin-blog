package models

type ArticleTag struct {
	ArticleModel Article `gorm:"foreignkey:ArticleID"`
	ArticleID    uint    `json:"article_id"`
	TagModel     Tag     `gorm:"foreignkey:TagID"`
	TagID        uint    `json:"tag_id"`
}
