package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"gorm.io/gorm"
)

func (articleApi *ArticleApi) LikeArticleView(c *gin.Context) {
	articleId := c.Param("id")
	err := global.DB.Model(&models.Article{}).Where("id=?", articleId).
		UpdateColumn("digg_count", gorm.Expr("digg_count+1")).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("点赞失败", c)
		return
	}
	res.OKWithMsg("点赞成功", c)
}
