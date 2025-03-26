package userApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/models/res"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	Nickname string     `json:"nickname"`
	UserID   uint       `json:"user_id" binding:"required" msg:"用户ID错误"`
}

// UserUpdateRoleView 用户权限变更,可以修改昵称
func (userApi *UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.User
	err = global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":     cr.Role,
		"nickname": cr.Nickname,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改权限失败", c)
		return
	}
	res.OKWithMsg("修改权限成功", c)
}
