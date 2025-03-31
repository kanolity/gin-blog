package commentApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"gorm.io/gorm"
)

func (commentApi *CommentApi) CommentLikeView(c *gin.Context) {
	commentID := c.Param("id")
	err := global.DB.Model(&models.Comment{}).
		Where("id = ?", commentID).
		UpdateColumn("digg_count", gorm.Expr("digg_count + 1")).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("点赞失败，请稍后再试", c)
		return
	}

	res.OKWithMsg("点赞成功", c)
}
