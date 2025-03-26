package menuApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/models/res"
)

type ImageSort struct {
	ImageId uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                //切换的时间,单位:s
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                    //切换的时间,单位:s
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` //菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                          //具体图片的顺序
}

// MenuCreateView 创建菜单
// @Tags         菜单管理
// @Summary      创建菜单
// @Description  创建菜单
// @Param data body MenuRequest  true "菜单参数"
// @Router       /api/menus [post]
// @Produce json
// @Success      200  {object}  res.Resp{data=string}
// @Failure      200  {object}  res.Resp{data=string}
func (menuApi *MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
	}

	var menu models.Menu
	err = global.DB.Take(&menu, "title=? or path=?", cr.Title, cr.Path).Error
	if err == nil {
		res.FailWithMsg("重复的菜单", c)
		return
	}

	//创建banner数据入库
	menuModel := models.Menu{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OKWithMsg("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBanner
	for _, sort := range cr.ImageSortList {
		//需判断image_id是否有这张图片
		menuBannerList = append(menuBannerList, models.MenuBanner{
			MenuID:   menuModel.ID,
			Sort:     sort.Sort,
			BannerID: sort.ImageId,
		})
	}
	//给第三张表入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单图片关联失败", c)
		return
	}
	res.OKWithMsg("菜单添加成功", c)
}
