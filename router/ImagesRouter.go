package router

import (
	"go_code/gin-vue-blog/api"
)

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("/image", imagesApi.ImageUploadView)
	router.GET("/image", imagesApi.ImageListView)
}
