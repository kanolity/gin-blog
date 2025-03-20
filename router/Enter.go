package router

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	routerGroupApi := RouterGroup{apiRouterGroup}
	//系统配置api
	routerGroupApi.InitSettingApiRouter()

	return router
}
