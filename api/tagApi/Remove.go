package tagApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

// TagRemoveView 批量删除标签
// @Tags         标签管理
// @Summary      批量删除标签
// @Description  批量删除标签
// @Param  data query  uint true "标签id"
// @Router       /api/adverts [delete]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
func (tagApi *TagApi) TagRemoveView(c *gin.Context) {
	id := c.Query("id")

	err := global.DB.Where("id=?", id).Delete(&models.Tag{}).Error
	if err != nil {
		global.Log.Error(err)
		res.OKWithMsg("删除标签失败", c)
		return
	}

	res.OKWithMsg("删除标签成功", c)
}
