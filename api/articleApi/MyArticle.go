package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service/common"
	"go_code/gin-vue-blog/utils/jwts"
)

// MyArticleListView 用户自己发布的文章列表
func (articleApi *ArticleApi) MyArticleListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.Article{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	articleList := make([]ArticleListResponse, 0, len(list))
	for _, article := range list {
		articleList = append(articleList, ArticleListResponse{
			ID:         article.ID,
			Title:      article.Title,
			NickName:   article.NickName,
			BannerPath: article.BannerPath,
			Tags:       article.Tags,
		})
	}
	res.OKWithList(articleList, count, c)
}
