package router

import "go_blog/api"

func (r RouterGroup) UserRouter() {
	UserApi := api.ApiGroupApp.UserApi
	r.POST("/email_login", UserApi.EmailLoginView)
}
