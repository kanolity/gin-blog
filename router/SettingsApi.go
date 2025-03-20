package router

import (
	"go_code/gin-vue-blog/api/settingsApi"
)

func (router RouterGroup) InitSettingApiRouter() {
	settingsApiRouter := settingsApi.SettingsApi{}
	router.GET("/settings/:name", settingsApiRouter.SettingsInfoView)
	router.PUT("/settings/:name", settingsApiRouter.SettingsInfoUpdateView)
}
