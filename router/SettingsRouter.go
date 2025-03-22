package router

import (
	"go_code/gin-vue-blog/api"
)

func (router RouterGroup) InitSettingApiRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/settings/:name", settingsApi.SettingsInfoView)
	router.PUT("/settings/:name", settingsApi.SettingsInfoUpdateView)
}
