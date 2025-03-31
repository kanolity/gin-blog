package menuApi

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"gorm.io/gorm"
)

func (menuApi *MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	id := c.Param("id")
	var menuModel models.Menu
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 清空旧关联
		err = tx.Model(&menuModel).Association("Banners").Clear()
		if err != nil {
			global.Log.Error("清空 Banner 关联失败：", err)
			return err
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
			err = tx.Create(&bannerList).Error
			if err != nil {
				global.Log.Error("创建菜单图片失败:", err)
				return err
			}
		}

		maps := structs.Map(&cr)
		err = tx.Model(&menuModel).Updates(maps).Error
		if err != nil {
			global.Log.Error("修改菜单失败:", err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改菜单失败", c)
		return
	}
	res.OKWithMsg("修改菜单成功", c)
}
