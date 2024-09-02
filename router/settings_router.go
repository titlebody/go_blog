package router

import (
	"go_blog/api"
)

func (r routerGroup) SettingsRouter() {
	SettingsApi := api.ApiGroupApp.SettingsApi

	r.GET("/settings", SettingsApi.SettingsInfoView)
	r.PUT("/settings", SettingsApi.SettingsUpdateView)
}
