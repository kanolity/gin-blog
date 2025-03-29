package userApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"gorm.io/gorm"
)

func (userApi *UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var userList []models.User
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("用户不存在", c)
		return
	}
	//事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//TODO:删除用户，消息表，评论表，发布、收藏的文章
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除用户失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("删除%d个用户", count), c)
}
