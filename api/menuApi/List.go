package menuApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.Menu
	Banners []Banner `json:"banners"`
}

func (menuApi *MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.Menu
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)

	var menuBanners []models.MenuBanner
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id IN (?)", menuIDList)
	var menus []MenuResponse

	for _, menu := range menuList {
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
		menus = append(menus, MenuResponse{
			Menu:    menu,
			Banners: banners,
		})
	}
	res.OKWithData(menus, c)
	return
}
