package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

//TODO 删除文章的所有评论和收藏其用户的收藏

// ArticleRemoveAdminView 管理员删除文章
func (articleApi *ArticleApi) ArticleRemoveAdminView(c *gin.Context) {
	id := c.Query("id")

	err := global.DB.Where("id=?", id).Delete(&models.Article{}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除文章失败", c)
		return
	}
	res.OKWithMsg("删除文章成功", c)
}
