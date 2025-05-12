package adApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

type AdRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        //标题
	Href   string ` json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`    //跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"` //图片
	IsShow bool   `json:"is_show" structs:"is_show"`                                   //是否展示
}

// AdCreateView 广告创建
// @Tags         广告管理
// @Summary      创建广告
// @Description  创建广告
// @Param  data body AdRequest true "表示多个参数"
// @Router       /api/adverts [post]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
// @Failure      200  {object}  res.Resp{data=string}
// @Failure      200  {object}  res.Resp{data=string}
func (adApi *AdApi) AdCreateView(c *gin.Context) {
	var cr AdRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//是否重复
	var ad models.Ad
	err = global.DB.Take(&ad, "title=?", cr.Title).Error
	if err == nil {
		res.FailWithMsg("该广告已存在", c)
		return
	}

	err = global.DB.Create(&models.Ad{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加广告失败", c)
		return
	}
	res.OKWithMsg("添加广告成功", c)
}
