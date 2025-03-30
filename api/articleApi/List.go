package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service/common"
)

type ArticleListResponse struct {
	ID         uint        `json:"id"`
	Title      string      `json:"title"`
	NickName   string      `json:"nick_name"`
	BannerPath string      `json:"banner_path"`
	Tags       ctype.Array `json:"tags"`
}

// ArticleListView 文章列表
func (articleApi *ArticleApi) ArticleListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.Article{}, common.Option{
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
