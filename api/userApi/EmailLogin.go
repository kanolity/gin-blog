package userApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
	"go_code/gin-vue-blog/utils/pwd"
)

type EmailLoginRequest struct {
	Username string `json:"username" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (userApi *UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var userModel models.User
	err = global.DB.Take(&userModel, "username=? or email=?", cr.Username, cr.Username).Error
	if err != nil {
		global.Log.Warn("用户名不存在")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   userModel.ID,
		Nickname: userModel.Nickname,
		Role:     int(userModel.Role),
	})
	if err != nil {
		res.FailWithMsg("token生成失败", c)
		return
	}
	res.OKWithData(token, c)
}
