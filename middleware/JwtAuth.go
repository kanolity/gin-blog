package middleware

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service"
	"go_code/gin-vue-blog/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}

		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}
		//判断是否在redis中
		if service.ServiceGroupApp.RedisService.CheckLogout(token) {
			res.FailWithMsg("token已失效", c)
			c.Abort()
			return
		}
		//登录的用户
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}

		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}
		//判断是否在redis中
		if service.ServiceGroupApp.RedisService.CheckLogout(token) {
			res.FailWithMsg("token已失效", c)
			c.Abort()
			return
		}
		//判断权限
		if ctype.Role(claims.Role) != ctype.Admin {
			res.FailWithMsg("权限错误", c)
			c.Abort()
			return
		}

		//登录的用户
		c.Set("claims", claims)
	}
}
