package router

import (
	"go_blog/api"
)

func (r RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	r.POST("/menus", app.MenuCreateView)
	r.GET("/menus", app.MenuListView)
	r.GET("/menu_names", app.MenuNameList)
	r.PUT("/menus/:id", app.MenuUpdateView)
	r.DELETE("/menus", app.MenuRemoveView)
	r.GET("/menus/:id", app.MenuDetailView)

}
