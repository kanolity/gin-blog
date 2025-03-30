package articleApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
	"gorm.io/gorm"
)

// ArticleRemoveUserView 用户删除自己发布的文章
func (articleApi *ArticleApi) ArticleRemoveUserView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var count int64
	global.DB.Model(&models.Article{}).Where("id IN ? AND user_id = ?", cr.IDList, claims.UserID).Count(&count)
	if count == 0 {
		res.FailWithMsg("文章不存在或您没有权限删除这些文章", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Where("id IN ? and user_id in ?", cr.IDList, claims.UserID).Delete(&models.Article{}).Error
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
