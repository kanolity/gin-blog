package imagesApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名"`
}

// ImageUpdateView 修改图片名称
// @Tags         图片管理
// @Summary      修改图片名称
// @Description  修改图片名称
// @Param  data body ImageUpdateRequest true "图片的部分参数"
// @Router       /api/images [put]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
func (imagesApi *ImagesApi) ImageUpdateView(c *gin.Context) {
	var req ImageUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}
	var image models.Banner
	err = global.DB.Take(&image, req.ID).Error
	if err != nil {
		res.FailWithMsg("文件不存在", c)
		return
	}
	err = global.DB.Model(&image).Update("name", req.Name).Error
	if err != nil {
		res.OKWithMsg(err.Error(), c)
		return
	}
	res.OKWithMsg("图片名称修改成功", c)
	return
}
