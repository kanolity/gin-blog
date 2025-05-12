package adApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

// AdRemoveView 删除广告
// @Tags         广告管理
// @Summary      删除广告
// @Description  删除广告
// @Param  data query  uint true "广告id列表"
// @Router       /api/adverts [delete]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
func (adApi *AdApi) AdRemoveView(c *gin.Context) {
	id := c.Query("id")
	err := global.DB.Where("id=?", id).Delete(&models.Ad{}).Error
	if err != nil {
		global.Log.Error(err)
		res.OKWithMsg("删除广告失败", c)
		return
	}
	res.OKWithMsg("删除广告成功", c)
}
