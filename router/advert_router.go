package router

import "go_blog/api"

func (r RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi

	r.POST("/adverts", app.AdvertCreateView)

}
