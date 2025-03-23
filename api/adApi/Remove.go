package adApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

// AdRemoveView 批量删除广告
// @Tags         广告管理
// @Summary      批量删除广告
// @Description  批量删除广告
// @Param  data body  models.RemoveRequest true "广告id列表"
// @Router       /api/adverts [delete]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
func (adApi *AdApi) AdRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var adList []models.Ad
	count := global.DB.Model(&models.Ad{}).Find(&adList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}
	global.DB.Delete(&adList)
	res.OKWithMsg(fmt.Sprintf("删除%d个广告", count), c)
}
