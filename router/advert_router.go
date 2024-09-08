package router

import "go_blog/api"

func (r RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi

	r.POST("/adverts", app.AdvertCreateView)
	r.GET("/adverts", app.AdvertListView)
	r.PUT("/adverts/:id", app.AdvertUpdateView)
	r.DELETE("/adverts", app.AdvertRemoveView)

}
