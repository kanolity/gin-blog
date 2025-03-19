package models

type Tag struct {
	Model
	Title    string    `gorm:"size:10" json:"title"`
	Articles []Article `gorm:"many2many:article_tag" json:"-"`
}
