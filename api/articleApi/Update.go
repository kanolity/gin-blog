package articleApi

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"gorm.io/gorm"
)

// ArticleUpdateView 修改文章
func (articleApi *ArticleApi) ArticleUpdateView(c *gin.Context) {
	//绑定参数
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")

	//获取文章和封面信息
	var articleModel models.Article
	err = global.DB.Take(&articleModel, id).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	var banner models.Banner
	err = global.DB.Take(&banner, cr.BannerID).Error
	if err != nil {
		res.FailWithMsg("文章封面不存在", c)
		return
	}
	maps := structs.Map(&cr)
	maps["banner_path"] = banner.Path

	var tagList []models.Tag
	for _, tag := range cr.Tags {
		var tagModel models.Tag
		err = global.DB.Take(&tagModel, "title=?", tag).Error
		if err != nil {
			res.FailWithMsg(fmt.Sprintf("标签不存在：%s", tag), c)
			return
		}
		tagList = append(tagList, tagModel)

	}

	//更改信息
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 清除旧的标签关联
		err = tx.Model(&articleModel).Association("TagModels").Clear()
		if err != nil {
			global.Log.Error("清除旧标签关联失败：", err)
			return err
		}

		// 插入新的标签关联（如果有标签列表）
		if len(tagList) > 0 {
			tagAssociations := make([]models.ArticleTag, 0, len(tagList))
			for _, tag := range tagList {
				tagAssociations = append(tagAssociations, models.ArticleTag{
					ArticleID: articleModel.ID,
					TagID:     tag.ID,
				})
			}
			if err = tx.Create(&tagAssociations).Error; err != nil {
				global.Log.Error("创建标签关联失败：", err)
				return err
			}
		}
		err = tx.Model(&articleModel).Updates(maps).Error
		if err != nil {
			global.Log.Error("修改文章失败:", err)
			return err
		}
		return nil
	})

	if err != nil {
		res.FailWithMsg("修改文章失败", c)
		return
	}
	res.OKWithMsg("修改文章成功", c)
}
