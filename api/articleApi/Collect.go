package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func (articleApi *ArticleApi) CollectArticleView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	articleId, _ := strconv.Atoi(c.Param("id"))

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		var existingCollect models.UserCollects
		err := tx.Where("user_id = ? AND article_id = ?", claims.UserID, uint(articleId)).Take(&existingCollect).Error
		if err == nil {
			res.FailWithMsg("您已收藏过该文章", c)
			return nil
		}

		//创建关联表
		err = tx.Create(&models.UserCollects{
			UserID:     claims.UserID,
			ArticleID:  uint(articleId),
			CreateTime: time.Now(),
		}).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		//收藏量增加
		err = tx.Model(&models.Article{}).Where("id=?", uint(articleId)).
			UpdateColumn("collects_count", gorm.Expr("collects_count+1")).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})

	if err != nil {
		res.FailWithMsg("收藏失败", c)
		return
	}
	res.OKWithMsg("收藏成功", c)
}
