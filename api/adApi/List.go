package adApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service/common"
)

// AdListView 广告列表
// @Tags         广告管理
// @Summary      广告列表
// @Description  广告列表
// @Param  data query models.PageInfo false "查询参数"
// @Router       /api/adverts [get]
// @Produce json
// @Success      200  {object}  res.Resp{data=res.ListResp[models.Ad]}
func (adApi *AdApi) AdListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//判断referer是否包含admin
	list, count, _ := common.ComList(models.Ad{IsShow: true}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	fmt.Println(list)
	res.OKWithList(list, count, c)
}
