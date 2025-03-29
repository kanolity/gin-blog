package router

import (
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

func (router RouterGroup) SettingRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings/:name", middleware.JwtAuth(), settingsApi.SettingsInfoView)
	router.PUT("/settings/:name", middleware.JwtAuth(), settingsApi.SettingsInfoUpdateView)
}
