package menuApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
)

func (menuApi *MenuApi) MenuRemoveView(c *gin.Context) {
	id := c.Query("id")
	var count int64
	global.DB.Model(&models.Menu{}).Where("id = ?", id).Count(&count)
	if count == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}

	err := global.DB.Where("id = ?", id).Delete(&models.Menu{}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除菜单失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("删除%d个菜单", count), c)

}
