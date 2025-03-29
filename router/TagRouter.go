package router

import (
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

func (router RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	router.POST("/tags", middleware.JwtAdmin(), tagApi.TagCreateView)
	router.GET("/tags", middleware.JwtAdmin(), tagApi.TagListView)
	router.PUT("/tags/:id", middleware.JwtAdmin(), tagApi.TagUpdateView)
	router.DELETE("/tags", middleware.JwtAdmin(), tagApi.TagRemoveView)
}
