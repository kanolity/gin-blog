package menuApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"gorm.io/gorm"
)

func (menuApi *MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var count int64
	global.DB.Model(&models.Menu{}).Where("id IN ?", cr.IDList).Count(&count)
	if count == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}

	//事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Where("id IN ?", cr.IDList).Delete(&models.Menu{}).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除菜单失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("删除%d个菜单", count), c)

}
