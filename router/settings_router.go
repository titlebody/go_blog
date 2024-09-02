package router

import (
	"go_blog/api"
)

func (r routerGroup) SettingsRouter() {
	SettingsApi := api.ApiGroupApp.SettingsApi

	r.GET("/settings/:name", SettingsApi.SettingsInfoView)
	r.PUT("/settings/:name", SettingsApi.SettingsUpdateView)
}
