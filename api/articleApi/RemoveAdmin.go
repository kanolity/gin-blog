package articleApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"gorm.io/gorm"
)

//TODO 删除文章的所有评论和收藏其用户的收藏

// ArticleRemoveAdminView 管理员批量删除文章
func (articleApi *ArticleApi) ArticleRemoveAdminView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var count int64
	global.DB.Model(&models.Article{}).Where("id IN ? ", cr.IDList).Count(&count)
	if count == 0 {
		res.FailWithMsg("文章不存在", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Where("id IN ?", cr.IDList).Delete(&models.Article{}).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除文章失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("删除%d个文章", count), c)
}
