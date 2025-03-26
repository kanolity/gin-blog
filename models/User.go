package models

import "go_code/gin-vue-blog/models/ctype"

// User 用户表
type User struct {
	Model
	Nickname       string           `gorm:"size:36" json:"nickname"`                                                       //昵称
	Username       string           `gorm:"size:36" json:"username"`                                                       //用户名
	Password       string           `gorm:"size:255" json:"-"`                                                             //密码
	Avatar         string           `gorm:"size:255" json:"avatar"`                                                        //头像
	Email          string           `gorm:"size:255" json:"email"`                                                         //邮箱
	Phone          string           `gorm:"size:20" json:"phone"`                                                          //电话
	Address        string           `gorm:"size:255" json:"address"`                                                       //地址
	Token          string           `gorm:"size:255" json:"token"`                                                         //其他平台的唯一ID
	IP             string           `gorm:"size:255" json:"ip"`                                                            //ip
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`                                                  //权限
	SignStatus     ctype.SignStatus `gorm:"default:1" json:"sign_status"`                                                  //注册来源
	ArticleModels  []Article        `gorm:"foreignkey:UserID" json:"-"`                                                    //发布的文章列表
	CollectsModels []Article        `gorm:"many2many:user_collects;jonForeignKey:UserID;JoinReference:ArticleID" json:"-"` //收藏的文章列表
}
