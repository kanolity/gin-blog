package adApi

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

// AdUpdateView 更新广告
// @Tags         广告管理
// @Summary      更新广告
// @Description  更新广告
// @Param  data body  AdRequest true "广告的参数"
// @Param id path int true "广告id"
// @Router       /api/adverts/{id} [put]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
func (adApi *AdApi) AdUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//标题是否重复
	var ad models.Ad
	err = global.DB.Take(&ad, "title=? AND id!=?", cr.Title, id).Error
	if err == nil {
		res.FailWithMsg("广告标题已存在", c)
		return
	}

	err = global.DB.Take(&ad, id).Error
	if err != nil {
		res.FailWithMsg("广告不存在", c)
		return
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&ad).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改广告失败", c)
		return
	}
	res.OKWithMsg("修改广告成功", c)
}
