package router

import (
	"go_code/gin-vue-blog/api"
)

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("/images", imagesApi.ImageUploadView)
	router.GET("/images", imagesApi.ImageListView)
	router.DELETE("/images", imagesApi.ImageRemoveView)
	router.PUT("/images", imagesApi.ImageUpdateView)
}
