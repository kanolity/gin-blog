package router

import "go_code/gin-vue-blog/api"

func (router RouterGroup) AdRouter() {
	adApi := api.ApiGroupApp.AdApi
	router.POST("/adverts", adApi.AdCreateView)
	router.GET("/adverts", adApi.AdListView)
	router.PUT("/adverts/:id", adApi.AdUpdateView)
	router.DELETE("/adverts", adApi.AdRemoveView)
}
