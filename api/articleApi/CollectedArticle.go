package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
)

func (articleApi *ArticleApi) CollectedArticleView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var articleList []ArticleListResponse
	err := global.DB.Table("articles").Joins("left join user_collects on user_collects.article_id = articles.id").
		Select("id", "title", "nick_name", "banner_path", "tags").
		Where("user_collects.user_id=?", claims.UserID).Find(&articleList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询失败", c)
		return
	}
	count := len(articleList)
	if count == 0 {
		res.OKWithMsg("未收藏文章", c)
		return
	}
	res.OKWithList(articleList, int64(count), c)
}
