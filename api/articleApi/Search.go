package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models/res"
)

type SearchArticleResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Abstract   string `json:"abstract"`
	NickName   string `json:"nick_name"`
	BannerPath string `json:"banner_path"`
}

func (articleApi *ArticleApi) SearchArticleView(c *gin.Context) {
	var articles []SearchArticleResponse
	category := c.Query("category")
	tagId := c.Query("tag_id")

	if category != "" {
		global.DB.Select("id", "title", "abstract", "nick_name", "banner_path").
			Where("category = ?", category).Find(&articles)

	}
	if tagId != "" {
		global.DB.Table("articles").Joins("left join article_tags on article_tags.article_id=articles.id").
			Select("id", "title", "abstract", "nick_name", "banner_path").
			Where("article_tags.tag_id=?", tagId).Find(&articles)
	}
	if len(articles) == 0 {
		res.OKWithMsg("没有相关文章", c)
		return
	}
	res.OKWithData(articles, c)
}
