package router

import (
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

func (router RouterGroup) MessageRouter() {
	msgApi := api.ApiGroupApp.MessageApi
	router.POST("/messages", middleware.JwtAuth(), msgApi.MessageCreateView)
	router.GET("/messages_all", middleware.JwtAdmin(), msgApi.MessageListAllView)
	router.GET("/messages", middleware.JwtAuth(), msgApi.MessageListView)
	router.GET("/chat_log", middleware.JwtAuth(), msgApi.ChatLogView)
}
