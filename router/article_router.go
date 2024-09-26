package router

import (
	"go_blog/api"
	"go_blog/middleware"
)

func (r RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	r.POST("/articles", middleware.JwtAuth(), app.ArticleCreateView)
	r.DELETE("/articles", middleware.JwtAuth(), app.ArticleRemoveView)
	r.PUT("/articles/:id", middleware.JwtAuth(), app.ArticleUpdateView)
	r.GET("/articles_list", app.ArticleListView)
	r.GET("/articles/:id", app.ArticleView)
	r.GET("/articles_search", app.ArticleSearchView)

}
