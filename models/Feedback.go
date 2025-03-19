package models

type Feedback struct {
	Model
	Email        string `gorm:"size:255" json:"email"`
	Content      string `gorm:"size:255" json:"content"`
	ApplyContent string `gorm:"size:255" json:"apply_content"` //回复的内容
	IsApply      bool   `json:"is_apply"`                      //是否回复
}
