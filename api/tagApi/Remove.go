package tagApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

// TagRemoveView 批量删除标签
// @Tags         标签管理
// @Summary      批量删除标签
// @Description  批量删除标签
// @Param  data body  models.RemoveRequest true "标签id列表"
// @Router       /api/adverts [delete]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
func (tagApi *TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var count int64
	global.DB.Model(&models.Tag{}).Where("id IN ?", cr.IDList).Count(&count)
	if count == 0 {
		res.FailWithMsg("标签不存在", c)
		return
	}

	err = global.DB.Where("id IN ?", cr.IDList).Delete(&models.Tag{}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除标签失败", c)
		return
	}

	res.OKWithMsg(fmt.Sprintf("删除%d个标签", count), c)
}
