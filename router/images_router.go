package router

import "go_blog/api"

func (r RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi

	r.POST("/images", app.ImageUploadView)
	r.GET("/images", app.ImageListView)
	r.GET("/images_name", app.ImageNameListView)
	r.DELETE("/images", app.ImageRemoveView)
	r.PUT("/images", app.ImageUpdateView)
}
