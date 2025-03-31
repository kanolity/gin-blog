package router

import (
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

func (router RouterGroup) CommentRouter() {
	commentApi := api.ApiGroupApp.CommentApi
	router.POST("comments", middleware.JwtAuth(), commentApi.CreateCommentView)
	router.GET("comments/:article_id", middleware.JwtAuth(), commentApi.CommentListView)
	router.PUT("comments/:id", middleware.JwtAuth(), commentApi.CommentLikeView)
	router.DELETE("comments/:id", middleware.JwtAuth(), commentApi.CommentRemoveView)
}
