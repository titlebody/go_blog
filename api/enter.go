package api

import (
	"go_blog/api/advert_api"
	"go_blog/api/images_api"
	"go_blog/api/menus_api"
	"go_blog/api/settings_api"
	"go_blog/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menus_api.MenuApi
	UserApi     user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
