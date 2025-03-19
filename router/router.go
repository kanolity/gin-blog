package router

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/api/settingsApi"
	"go_code/gin-vue-blog/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	//api
	settingsApi := settingsApi.SettingsApi{}
	ApiRouterGroup := router.Group("api")
	ApiRouterGroup.GET("/settings", settingsApi.SettingsInfoView)

	return router
}
