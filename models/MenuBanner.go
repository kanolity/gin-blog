package models

// MenuImageModel 自定义菜单和背景图的连接表，方便排序
type MenuImageModel struct {
	MenuID      uint   `json:"menu_id"`             // 菜单 ID
	MenuModel   Menu   `gorm:"foreignKey:MenuID"`   // 菜单模型
	BannerID    uint   `json:"banner_id"`           // 图片 ID
	BannerModel Banner `gorm:"foreignKey:BannerID"` // 图片模型
	Sort        int    `gorm:"size:10" json:"sort"` // 排序字段
}
