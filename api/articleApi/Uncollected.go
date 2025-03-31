package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
	"gorm.io/gorm"
	"strconv"
)

func (articleApi *ArticleApi) UncollectedArticleView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	articleId, _ := strconv.Atoi(c.Param("id"))

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("user_id=? and article_id=?", claims.UserID, uint(articleId)).
			Delete(&models.UserCollects{}).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}

		//收藏量减少
		err = tx.Model(&models.Article{}).Where("id=?", uint(articleId)).
			UpdateColumn("collects_count", gorm.Expr("collects_count-1")).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("取消收藏失败", c)
		return
	}
	res.OKWithMsg("取消收藏成功", c)
}
