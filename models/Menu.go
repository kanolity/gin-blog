package models

import "go_code/gin-vue-blog/models/ctype"

// Menu 菜单表
type Menu struct {
	Model
	MenuTitle    string      `gorm:"size:32" json:"menu_title"`                                                               // 菜单标题
	MenuTitleEn  string      `gorm:"size:32" json:"menu_title_en"`                                                            // 英文标题
	Slogan       string      `gorm:"size:64" json:"slogan"`                                                                   // 宣传标语
	Abstract     ctype.Array `gorm:"type:string" json:"abstract"`                                                             // 简介
	AbstractTime int         `json:"abstract_time"`                                                                           // 简介切换时间
	Banners      []Banner    `gorm:"many2many:menu_banner;joinForeignKey:MenuID;JoinReferences:bannerID" json:"menu_banners"` // 菜单的图片列表
	BannerTime   int         `json:"menu_time"`                                                                               // 菜单图片切换时间，为0表示不切换
	Sort         int         `gorm:"size:10" json:"sort"`                                                                     // 菜单的排序顺序
}
