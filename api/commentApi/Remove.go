package commentApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"gorm.io/gorm"
)

func (commentApi *CommentApi) CommentRemoveView(c *gin.Context) {
	commentID := c.Param("id")

	// 查询评论是否存在
	var comment models.Comment
	err := global.DB.Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除评论失败", c)
		return
	}

	// 递归删除评论及其子评论
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = deleteSubComments(tx, commentID, comment.ArticleID)
		if err != nil {
			return err
		}

		// 删除评论本身
		err = tx.Delete(&models.Comment{}, commentID).Error
		if err != nil {
			global.Log.Error("删除评论失败：", err)
			return err
		}

		// 如果是子评论，更新父评论的 CommentCount
		if comment.ParentCommentID != nil {
			err = tx.Model(&models.Comment{}).
				Where("id = ?", comment.ParentCommentID).
				UpdateColumn("comment_count", gorm.Expr("comment_count - 1")).Error
			if err != nil {
				return err
			}
		}

		// 更新文章的评论数
		err = tx.Model(&models.Article{}).
			Where("id = ?", comment.ArticleID).
			UpdateColumn("comment_count", gorm.Expr("comment_count - 1")).Error
		if err != nil {
			global.Log.Error("更新文章评论数失败：", err)
			return err
		}
		return nil
	})
	if err != nil {
		res.FailWithMsg("删除评论失败", c)
		return
	}

	res.OKWithMsg("评论删除成功", c)
}

// 递归删除子评论
func deleteSubComments(tx *gorm.DB, commentID string, articleID uint) error {
	var subComments []models.Comment
	err := tx.Where("parent_comment_id = ?", commentID).Find(&subComments).Error
	if err != nil {
		return err
	}

	for _, subComment := range subComments {
		err = deleteSubComments(tx, fmt.Sprintf("%d", subComment.ID), articleID)
		if err != nil {
			return err
		}
		// 每删除一个子评论，减少文章评论数
		err = tx.Model(&models.Article{}).
			Where("id = ?", articleID).
			UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			return err
		}
	}

	err = tx.Where("parent_comment_id = ?", commentID).Delete(&models.Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}
