package models

type Ad struct {
	Model
	Title  string `gorm:"size:32" json:"title"` //标题
	Href   string ` json:"href"`                //跳转链接
	Images string `json:"images"`               //图片
	IsShow bool   `json:"is_show"`              //是否展示
}
