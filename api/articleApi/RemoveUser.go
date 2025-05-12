package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
)

// ArticleRemoveUserView 用户删除自己发布的文章
func (articleApi *ArticleApi) ArticleRemoveUserView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	id := c.Query("id")

	err := global.DB.Model(&models.Article{}).Where("id = ? AND user_id = ?", id, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("文章不存在或您没有权限删除这些文章", c)
		return
	}

	err = global.DB.Where("id = ? and user_id in ?", id, claims.UserID).Delete(&models.Article{}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除文章失败", c)
		return
	}
	res.OKWithMsg("删除文章成功", c)
}
