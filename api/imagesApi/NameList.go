package imagesApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// ImageNameListView 图片名称列表
// @Tags         图片管理
// @Summary      图片名称列表
// @Description  图片名称列表
// @Router       /api/image_names [get]
// @Produce json
// @Success      200  {object}  res.Resp{data=[]ImageResponse}
func (imagesApi *ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse

	global.DB.Model(&models.Banner{}).Select("id", "path", "name").Scan(&imageList)

	res.OKWithData(imageList, c)
	return
}
