package userApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/pwd"
)

type UserCreateRequest struct {
	Nickname   string     `json:"nickname" binding:"required" msg:"请输入昵称"`  //昵称
	Username   string     `json:"username" binding:"required" msg:"请输入用户名"` //用户名
	Password   string     `json:"password" binding:"required" msg:"请输入密码"`  //密码
	RePassword string     `json:"re_password" binding:"required" msg:"请再次输入密码"`
	Role       ctype.Role `json:"role" binding:"required" msg:"请选择权限"` //权限

}

// UserCreateView 创建用户
func (userApi *UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//判断用户名是否存在
	var userModel models.User
	err = global.DB.Take(&userModel, "username=?", cr.Username).Error
	if err == nil {
		//存在
		res.FailWithMsg("用户名存在", c)
		return
	}

	if cr.Password != cr.RePassword {
		res.FailWithMsg("两次密码不一致", c)
		return
	}
	//hash密码
	hashPwd := pwd.HashPwd(cr.Password)

	//入库
	err = global.DB.Create(&models.User{
		Username:   cr.Username,
		Nickname:   cr.Nickname,
		Password:   hashPwd,
		Role:       cr.Role,
		SignStatus: ctype.Email,
		IP:         c.ClientIP(),
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("用户%s创建成功", cr.Username), c)
}
