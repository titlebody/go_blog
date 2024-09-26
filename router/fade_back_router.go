package router

import "go_blog/api"

func (r RouterGroup) FadeBackRouter() {
	app := api.ApiGroupApp.FadeBackApi
	r.POST("/fade_back", app.FadeBackView)
	r.GET("/fade_back_list", app.FadeBackListView)
	r.DELETE("/fade_back", app.RemoveFadeView)
}
