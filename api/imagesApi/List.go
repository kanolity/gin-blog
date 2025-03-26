package imagesApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service/common"
)

// ImageListView 图片列表
// @Tags         图片管理
// @Summary      图片列表
// @Description  图片列表
// @Param  data query models.PageInfo true "查询参数"
// @Router       /api/images [get]
// @Produce json
// @Success      200  {object}  res.Resp{data=res.ListResp[models.Banner]}
func (imagesApi *ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.Banner{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	if err != nil {
		res.FailWithMsg(err.Error(), c)
	}

	res.OKWithList(list, count, c)
	return
}
