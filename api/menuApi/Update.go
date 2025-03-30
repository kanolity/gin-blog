package menuApi

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

func (menuApi *MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
	}

	id := c.Param("id")
	var menuModel models.Menu
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}

	if len(cr.ImageSortList) > 0 {
		var bannerList []models.MenuBanner
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBanner{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageId,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("创建菜单图片失败", c)
			return
		}
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改菜单失败", c)
		return
	}
	res.OKWithMsg("修改菜单成功", c)
}
