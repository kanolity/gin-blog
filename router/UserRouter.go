package router

import (
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/email_login", userApi.EmailLoginView)
	router.GET("/users", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("/user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.PUT("/user_password", middleware.JwtAuth(), userApi.UserUpdatePasswordView)
}
