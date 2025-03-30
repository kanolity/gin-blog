package articleApi

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

// ArticleUpdateView 修改文章
func (articleApi *ArticleApi) ArticleUpdateView(c *gin.Context) {
	//绑定参数
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")

	//获取文章和封面信息
	var articleModel models.Article
	err = global.DB.Take(&articleModel, id).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	var banner models.Banner
	err = global.DB.Take(&banner, cr.BannerID).Error
	if err != nil {
		res.FailWithMsg("文章封面不存在", c)
		return
	}
	maps := structs.Map(&cr)
	maps["banner_path"] = banner.Path

	//更改信息
	err = global.DB.Model(&articleModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改文章失败", c)
		return
	}

	res.OKWithMsg("修改文章成功", c)
}
