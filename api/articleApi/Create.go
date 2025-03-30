package articleApi

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
	"gorm.io/gorm"
	"html"
)

type ArticleRequest struct {
	Title    string      `json:"title" binding:"required"`     //标题
	Abstract string      `json:"abstract"`                     // 文章简介
	Content  string      `json:"content" binding:"required"`   // 文章内容
	Category string      `json:"category" binding:"required"`  // 文章分类
	Source   string      `json:"source"`                       // 文章来源
	Link     string      `json:"link"`                         // 文章来源链接
	BannerID uint        `json:"banner_id" binding:"required"` // 文章封面ID
	Tags     ctype.Array `json:"tags"`                         // 文章标签
}

func (articleApi *ArticleApi) ArticleCreateView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 防范 XSS: 对用户输入数据进行转义或过滤
	cr.Title = html.EscapeString(cr.Title)
	cr.Abstract = html.EscapeString(cr.Abstract)
	cr.Content = html.EscapeString(cr.Content)
	cr.Category = html.EscapeString(cr.Category)
	cr.Source = html.EscapeString(cr.Source)
	cr.Link = html.EscapeString(cr.Link)
	for i, tag := range cr.Tags {
		cr.Tags[i] = html.EscapeString(tag)
	}

	//事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//查找用户和图片相关信息
		var user models.User
		tx.Take(&user, claims.UserID)
		var banner models.Banner
		tx.Take(&banner, cr.BannerID)

		// 检查是否存在重复的文章
		var existingArticle models.Article
		err = tx.Where("title = ?", cr.Title).Take(&existingArticle).Error
		if err == nil {
			return fmt.Errorf("文章标题已存在，请更换标题")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果错误不是 ErrRecordNotFound，则表示查询出错
			return err
		}

		article := models.Article{
			Title:      cr.Title,
			Abstract:   cr.Abstract,
			Content:    cr.Content,
			Category:   cr.Category,
			Source:     cr.Source,
			Link:       cr.Link,
			BannerID:   cr.BannerID,
			Tags:       cr.Tags,
			UserID:     user.ID,
			NickName:   user.Nickname,
			BannerPath: banner.Path,
		}
		//保存文章
		err = tx.Create(&article).Error
		if err != nil {
			return err
		}
		var tag models.Tag

		//保存标签，建立关联关系
		for _, tagTitle := range cr.Tags {
			// 检查标签是否存在，不存在则创建
			err = tx.FirstOrCreate(&tag, models.Tag{Title: tagTitle}).Error
			if err != nil {
				return err
			}
			//关联标签和文章
			err = tx.Model(&article).Association("TagModels").Append(&tag)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(fmt.Sprintf("文章创建失败：%v", err), c)
		return
	}
	res.OKWithMsg("文章创建成功", c)
}
