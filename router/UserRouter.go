package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

var store = cookie.NewStore([]byte("FG2d3T7Dh7gf32r2H83HR"))

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("/email_login", userApi.EmailLoginView)
	router.GET("/users", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("/user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.PUT("/user_password", middleware.JwtAuth(), userApi.UserUpdatePasswordView)
	router.POST("/logout", middleware.JwtAuth(), userApi.UserLogoutView)
	router.POST("/bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)
}
