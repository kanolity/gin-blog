package commentApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"time"
)

type CommentListResponse struct {
	ID           uint                  `json:"id"`
	Content      string                `json:"content"`
	DiggCount    int                   `json:"digg_count"`
	CommentCount int                   `json:"comment_count"`
	CreatedAt    time.Time             `json:"created_at"`
	SubComments  []CommentListResponse `json:"sub_comments"`
}

func (commentApi *CommentApi) CommentListView(c *gin.Context) {
	articleId := c.Param("article_id")

	// 查询父评论
	var comments []models.Comment
	err := global.DB.Model(models.Comment{}).Where("article_id = ? AND parent_comment_id IS NULL", articleId).
		Order("created_at DESC"). // 按创建时间倒序
		Find(&comments).Error
	if err != nil {
		global.Log.Error("评论列表查询失败：", err)
		res.FailWithMsg("查询评论失败", c)
		return
	}

	commentList := make([]CommentListResponse, 0, len(comments))
	// 查询子评论（递归加载所有子评论）
	for _, comment := range comments {
		var subComments []models.Comment
		global.DB.Where("parent_comment_id = ?", comment.ID).Find(&subComments)

		// 构造子评论
		subCommentList := make([]CommentListResponse, 0, len(subComments))
		for _, sub := range subComments {
			subCommentList = append(subCommentList, CommentListResponse{
				ID:           sub.ID,
				Content:      sub.Content,
				DiggCount:    sub.DiggCount,
				CommentCount: sub.CommentCount,
				CreatedAt:    sub.CreatedAt,
			})
		}
		//构造父评论
		commentList = append(commentList, CommentListResponse{
			ID:           comment.ID,
			Content:      comment.Content,
			DiggCount:    comment.DiggCount,
			CommentCount: comment.CommentCount,
			CreatedAt:    comment.CreatedAt,
			SubComments:  subCommentList,
		})
	}
	count := int64(len(commentList))

	res.OKWithList(commentList, count, c)
}
