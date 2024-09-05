package router

import "go_blog/api"

func (r RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi

	r.POST("/images", app.ImageUploadView)
	r.GET("/images", app.ImageListView)
	r.DELETE("/images", app.ImageRemoveView)
}
