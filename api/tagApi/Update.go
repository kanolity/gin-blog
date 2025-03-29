package tagApi

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

// TagUpdateView 更新标签
// @Tags        标签管理
// @Summary      更新标签
// @Description  更新标签
// @Param  data body  AdRequest true "标签的参数"
// @Param id path int true "标签id"
// @Router       /api/adverts/{id} [put]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
func (tagApi *TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//标题是否重复
	var tag models.Tag
	err = global.DB.Take(&tag, "title=? AND id!=?", cr.Title, id).Error
	if err == nil {
		res.FailWithMsg("标签已存在", c)
		return
	}

	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMsg("标签不存在", c)
		return
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改标签失败", c)
		return
	}
	res.OKWithMsg("修改标签成功", c)
}
