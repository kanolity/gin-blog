package imagesApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service/common"
)

// ImageListView 图片列表
func (imagesApi *ImagesApi) ImageListView(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.Banner{}, common.Option{
		PageInfo: page,
		Debug:    true,
	})
	if err != nil {
		res.FailWithMsg(err.Error(), c)
	}

	res.OKWithList(list, count, c)
	return
}
