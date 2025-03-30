package userApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service/common"
	"go_code/gin-vue-blog/utils/desens"
	"go_code/gin-vue-blog/utils/jwts"
)

func (userApi *UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, _ := common.ComList(models.User{}, common.Option{
		PageInfo: page,
		Debug:    false,
	})
	users := make([]models.User, 0, len(list))
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.Admin {
			//非管理员
			user.Username = ""
		}
		//数据脱敏
		user.Phone = desens.DesensitizationPhone(user.Phone)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, user)
	}
	res.OKWithList(users, count, c)
}
