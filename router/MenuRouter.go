package router

import (
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

func (router RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	router.POST("/menus", middleware.JwtAdmin(), menuApi.MenuCreateView)
	router.GET("/menus", middleware.JwtAdmin(), menuApi.MenuListView)
	router.GET("/menu_names", middleware.JwtAdmin(), menuApi.MenuNameListView)
	router.PUT("/menus/:id", middleware.JwtAdmin(), menuApi.MenuUpdateView)
	router.DELETE("/menus", middleware.JwtAdmin(), menuApi.MenuRemoveView)
	router.GET("/menus/:id", middleware.JwtAdmin(), menuApi.MenuInfoView)
}
