package router

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/api/settings_api"
	"go_code/gin-vue-blog/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	settingsApi := settings_api.SettingsApi{}
	ApiRouterGroup := router.Group("api")
	ApiRouterGroup.GET("/settings", settingsApi.SettingsInfoView)

	return router
}
