package imagesApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

// ImageRemoveView 删除图片
// @Tags         图片管理
// @Summary      删除图片
// @Description  删除图片
// @Param  data body models.RemoveRequest true "图片id列表"
// @Router       /api/images [delete]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
func (imagesApi *ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []models.Banner
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("文件不存在", c)
	}
	global.DB.Delete(&imageList)
	res.OKWithMsg(fmt.Sprintf("删除%d张图片", count), c)
}
