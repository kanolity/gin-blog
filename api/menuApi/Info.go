package menuApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

func (menuApi *MenuApi) MenuInfoView(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu
	err := global.DB.Take(&menu, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	var menuBanners []models.MenuBanner
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id=?", id)

	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		if menu.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}

	menuResponse := MenuResponse{
		Menu:    menu,
		Banners: banners,
	}

	res.OKWithData(menuResponse, c)
	return
}
