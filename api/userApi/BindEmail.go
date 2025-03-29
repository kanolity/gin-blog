package userApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/plugins/email"
	"go_code/gin-vue-blog/service"
	"go_code/gin-vue-blog/utils/jwts"
	"go_code/gin-vue-blog/utils/random"
)

type BindEmailRequest struct {
	Email string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code  *string `json:"code"`
}

func (userApi *UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	if cr.Code == nil {
		//第一次，后台发验证码
		//生成四位验证码
		code := random.Code()

		//用redis存储邮箱和验证码
		err = service.ServiceGroupApp.RedisService.BindEmail(code, cr.Email)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("校验码存储错误", c)
			return
		}

		//发送验证码
		err = email.NewCode().Send(cr.Email, "验证码："+code)
		if err != nil {
			global.Log.Error(err)
		}
		res.FailWithMsg("验证码已发送", c)
		return
	}

	//校验邮箱和验证码
	if !service.ServiceGroupApp.RedisService.CheckBind(*cr.Code, cr.Email) {
		res.FailWithMsg("验证码或邮箱错误", c)
		return
	}
	//修改邮箱
	var user models.User
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"email": cr.Email,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("绑定失败", c)
		return
	}
	res.OKWithMsg("绑定成功", c)
	return
}
