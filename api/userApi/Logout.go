package userApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service"
	"go_code/gin-vue-blog/utils/jwts"
	"time"
)

func (userApi *UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")

	//计算距离现在的过期时间
	exp := claims.ExpiresAt
	now := time.Now()

	diff := exp.Time.Sub(now)
	err := service.ServiceGroupApp.RedisService.Logout(token, diff)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("注销失败", c)
		return
	}
	res.OKWithMsg("注销成功", c)
}
