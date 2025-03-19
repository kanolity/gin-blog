package models

type Comment struct {
	Model

	SubComments        []*Comment `gorm:"foreignkey:ParentCommentID" json:"sub_comments"`  // 子评论列表
	ParentCommentModel *Comment   `gorm:"foreignkey:ParentCommentID" json:"comment_model"` // 父级评论
	ParentCommentID    *uint      `json:"parent_comment_id"`                               // 父评论ID
	Content            string     `gorm:"size:256" json:"content"`                         // 评论内容
	DiggCount          int        `gorm:"size:8;default:0" json:"digg_count"`              // 点赞数
	CommentCount       int        `gorm:"size:8;default:0" json:"comment_count"`           // 子评论数
	Article            Article    `gorm:"foreignkey:ArticleID" json:"-"`                   // 关联的文章
	ArticleID          uint       `json:"article_id"`                                      // 文章ID
	User               User       `json:"user"`                                            // 关联的用户
	UserID             uint       `json:"user_id"`                                         // 评论的用户
}
