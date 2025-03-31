package models

import "go_code/gin-vue-blog/models/ctype"

// Article 文章表
type Article struct {
	Model
	Title         string      `gorm:"size:32" json:"title"`            //标题
	Abstract      string      `json:"abstract"`                        // 文章简介
	Content       string      `json:"content"`                         // 文章内容
	LookCount     int         `gorm:"default:0" json:"look_count"`     // 浏览量
	CommentCount  int         `gorm:"default:0" json:"comment_count"`  // 评论数
	DiggCount     int         `gorm:"default:0" json:"digg_count"`     // 点赞数
	CollectsCount int         `gorm:"default:0" json:"collects_count"` // 收藏量
	TagModels     []Tag       `gorm:"many2many:article_tags" json:"-"` // 文章标签
	CommentModel  []Comment   `gorm:"ForeignKey:ArticleID" json:"-"`   // 文章的评论列表
	UserModel     User        `gorm:"foreignKey:UserID" json:"-"`      // 文章作者
	UserID        uint        `json:"user_id"`                         // 用户ID
	Category      string      `gorm:"size:20" json:"category"`         // 文章分类
	Source        string      `json:"source"`                          // 文章来源
	Link          string      `json:"link"`                            // 文章来源链接
	Banner        Banner      `gorm:"ForeignKey:BannerID" json:"-"`    // 文章封面
	BannerID      uint        `json:"banner_id"`                       // 文章封面ID
	NickName      string      `gorm:"size:42" json:"nick_name"`        // 发布文章的用户昵称
	BannerPath    string      `json:"banner_path"`                     // 文章的封面路径
	Tags          ctype.Array `gorm:"type:string;size:64" json:"tags"` // 文章标签
}
