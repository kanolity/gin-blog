package tagApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service/common"
)

// TagListView 标签列表
// @Tags         标签管理
// @Summary      标签列表
// @Description  标签列表
// @Param  data query models.PageInfo false "查询参数"
// @Router       /api/adverts [get]
// @Produce json
// @Success      200  {object}  res.Resp{data=res.ListResp[models.Ad]}
func (tagApi *TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, _ := common.ComList(models.Tag{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	res.OKWithList(list, count, c)
}
