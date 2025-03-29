package router

import (
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("/images", middleware.JwtAuth(), imagesApi.ImageUploadView)
	router.GET("/images", middleware.JwtAuth(), imagesApi.ImageListView)
	router.GET("/image_names", middleware.JwtAuth(), imagesApi.ImageNameListView)
	router.DELETE("/images", middleware.JwtAuth(), imagesApi.ImageRemoveView)
	router.PUT("/images", middleware.JwtAuth(), imagesApi.ImageUpdateView)
}
