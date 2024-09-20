package router

import "go_blog/api"

func (r RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi

	r.POST("/tags", app.TagCreateView)
	r.GET("/tags", app.TagListView)
	r.PUT("/tags/:id", app.TagUpdateView)
	r.DELETE("/tags", app.TagRemoveView)

}
