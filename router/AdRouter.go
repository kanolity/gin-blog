package router

import (
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

func (router RouterGroup) AdRouter() {
	adApi := api.ApiGroupApp.AdApi
	router.POST("/adverts", middleware.JwtAdmin(), adApi.AdCreateView)
	router.GET("/adverts", middleware.JwtAdmin(), adApi.AdListView)
	router.PUT("/adverts/:id", middleware.JwtAdmin(), adApi.AdUpdateView)
	router.DELETE("/adverts", middleware.JwtAdmin(), adApi.AdRemoveView)
}
