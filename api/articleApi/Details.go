package articleApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/models/res"
)

type ArticleDetailsResponse struct {
	Title         string      `json:"title"`          //标题
	Abstract      string      `json:"abstract"`       // 文章简介
	Content       string      `json:"content"`        // 文章内容
	LookCount     int         `json:"look_count"`     // 浏览量
	CommentCount  int         `json:"comment_count"`  // 评论数
	DiggCount     int         `json:"digg_count"`     // 点赞数
	CollectsCount int         `json:"collects_count"` // 收藏量
	Category      string      `json:"category"`       // 文章分类
	Source        string      `json:"source"`         // 文章来源
	Link          string      `json:"link"`           // 文章来源链接
	NickName      string      `json:"nick_name"`      // 发布文章的用户昵称
	BannerPath    string      `json:"banner_path"`    // 文章的封面路径
	Tags          ctype.Array `json:"tags"`           // 文章标签
}

func (articleApi *ArticleApi) ArticleDetailsView(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	err := global.DB.Take(&article, "id=?", id).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	articleDetails := ArticleDetailsResponse{
		Title:         article.Title,
		Abstract:      article.Abstract,
		Content:       article.Content,
		LookCount:     article.LookCount,
		CommentCount:  article.CommentCount,
		DiggCount:     article.DiggCount,
		CollectsCount: article.CollectsCount,
		Category:      article.Category,
		Source:        article.Source,
		Link:          article.Link,
		NickName:      article.NickName,
		BannerPath:    article.BannerPath,
		Tags:          article.Tags,
	}
	res.OKWithData(articleDetails, c)
}
