package tagApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` //标签
}

// TagCreateView 标签创建
// @Tags         标签管理
// @Summary      创建标签
// @Description  创建标签
// @Param  data body AdRequest true "表示多个参数"
// @Router       /api/adverts [post]
// @Produce json
// @Success      200  {object}  res.Resp{}
func (tagApi *TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//是否重复
	var tag models.Tag
	err = global.DB.Take(&tag, "title=?", cr.Title).Error
	if err == nil {
		res.FailWithMsg("该标签已存在", c)
		return
	}

	err = global.DB.Create(&models.Tag{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加标签失败", c)
		return
	}
	res.OKWithMsg("添加标签成功", c)
}
