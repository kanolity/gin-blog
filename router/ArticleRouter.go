package router

import (
	"go_code/gin-vue-blog/api"
	"go_code/gin-vue-blog/middleware"
)

func (router RouterGroup) ArticleRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	router.POST("/articles", middleware.JwtAuth(), articleApi.ArticleCreateView)
	router.GET("/articles", middleware.JwtAuth(), articleApi.ArticleListView)
	router.GET("/articles/:id", middleware.JwtAuth(), articleApi.ArticleDetailsView)
	router.GET("/user_articles", middleware.JwtAuth(), articleApi.MyArticleListView)
	router.PUT("/articles/:id", middleware.JwtAuth(), articleApi.ArticleUpdateView)
	router.DELETE("/articles", middleware.JwtAdmin(), articleApi.ArticleRemoveAdminView)
	router.DELETE("/my_articles", middleware.JwtAuth(), articleApi.ArticleRemoveUserView)
	router.GET("/search_articles", middleware.JwtAuth(), articleApi.SearchArticleView)
	router.POST("/articles/:id", middleware.JwtAuth(), articleApi.CollectArticleView)
	router.GET("/my_collects", middleware.JwtAuth(), articleApi.CollectedArticleView)
	router.DELETE("/articles/:id", middleware.JwtAuth(), articleApi.UncollectedArticleView)
	router.PUT("/articles/:id/like", middleware.JwtAuth(), articleApi.LikeArticleView)
	router.PUT("/articles/:id/unlike", middleware.JwtAuth(), articleApi.UnlikeArticleView)
	//TODO 文章评论量
}
