package messageApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service/common"
)

func (messageApi *MessageApi) MessageListAllView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, _ := common.ComList(models.Tag{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	res.OKWithList(list, count, c)
}
