package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"go_code/gin-vue-blog/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("api")

	routerGroupApi := RouterGroup{apiRouterGroup}

	routerGroupApi.SettingRouter()
	routerGroupApi.ImagesRouter()
	routerGroupApi.AdRouter()
	routerGroupApi.MenuRouter()
	routerGroupApi.UserRouter()
	routerGroupApi.TagRouter()
	routerGroupApi.MessageRouter()

	return router
}
