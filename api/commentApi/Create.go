package commentApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
	"gorm.io/gorm"
)

type CommentCreateRequest struct {
	ParentCommentID *uint  `json:"parent_comment_id"`
	Content         string `json:"content" binding:"required" msg:"评论不能为空"`
	ArticleID       uint   `json:"article_id" binding:"required"`
}

func (commentApi *CommentApi) CreateCommentView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr CommentCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithCode(res.ArgumentError, c)
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}

	comment := models.Comment{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&comment).Error
		if err != nil {
			global.Log.Error(err)
			return fmt.Errorf("评论创建失败：%w", err)
		}
		// 如果是子评论，更新父评论的 CommentCount
		if comment.ParentCommentID != nil {
			err = tx.Model(&models.Comment{}).
				Where("id = ?", *comment.ParentCommentID).
				UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error
			if err != nil {
				global.Log.Error(err)
				return fmt.Errorf("父评论的评论数更新失败：%w", err)
			}
		}
		//更新文章评论数
		err = tx.Model(&models.Article{}).Where("id=?", cr.ArticleID).
			UpdateColumn("comment_count", gorm.Expr("comment_count+1")).Error
		if err != nil {
			global.Log.Error(err)
			return fmt.Errorf("文章评论数更新失败：%w", err)
		}
		return nil
	})
	if err != nil {
		res.FailWithMsg(fmt.Sprintf("评论创建失败:%v", err), c)
	}
	res.OKWithMsg("评论创建成功", c)
}
