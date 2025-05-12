package imagesApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

// ImageRemoveView 删除图片
// @Tags         图片管理
// @Summary      删除图片
// @Description  删除图片
// @Param  data query uint true "图片id"
// @Router       /api/images [delete]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
// @Failure      200  {object}  res.Resp{data=string}
func (imagesApi *ImagesApi) ImageRemoveView(c *gin.Context) {
	id := c.Query("id")

	err := global.DB.Where("id=?", id).Delete(&models.Banner{})
	if err != nil {
		res.FailWithMsg("删除图片失败", c)
	}
	res.OKWithMsg("删除图片成功", c)
}
