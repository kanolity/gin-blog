package userApi

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
)

type UserUpdateInfoRequest struct {
	Nickname string `json:"nickname"` //昵称
	Avatar   string `json:"avatar"`   //头像
	Phone    string `json:"phone"`    //电话
	Address  string `json:"address"`  //地址
}

// UserUpdateInfoView 修改用户信息
func (userApi *UserApi) UserUpdateInfoView(c *gin.Context) {
	//从token中获取ID
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr UserUpdateInfoRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var user models.User
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&user).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改信息失败", c)
		return
	}
	res.OKWithMsg("修改信息成功", c)
}
