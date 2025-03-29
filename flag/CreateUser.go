package flag

import (
	"fmt"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/utils/pwd"
)

func CreateUser(permission string) {
	//创建用户
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名:")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称:")
	fmt.Scan(&nickName)
	fmt.Printf("请输入密码:")
	fmt.Scan(&password)
	fmt.Printf("请确认密码:")
	fmt.Scan(&rePassword)
	fmt.Printf("请输入邮箱:")
	fmt.Scan(&email)

	//判断用户名是否存在
	var userModel models.User
	err := global.DB.Take(&userModel, "username=?", userName).Error
	if err == nil {
		//存在
		global.Log.Error("用户名存在，请重新输入")
		return
	}
	//校验密码
	if password != rePassword {
		global.Log.Error("密码不一致，请重新输入")
		return
	}
	role := ctype.NormalUser
	if permission == "admin" {
		role = ctype.Admin
	}

	//hash密码
	hashPwd := pwd.HashPwd(password)

	//头像
	avatar := "/uploads/file/avatar/UMP45.png"
	//入库
	err = global.DB.Create(&models.User{
		Username:   userName,
		Nickname:   nickName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		SignStatus: ctype.Email,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Address:    "地址",
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功", userName)
}
