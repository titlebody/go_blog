package router

import (
	"go_blog/api"
	"go_blog/middleware"
)

func (r RouterGroup) UserRouter() {
	UserApi := api.ApiGroupApp.UserAPI
	r.POST("/email_login", UserApi.EmailLoginView)
	r.PUT("/user_role", middleware.JwtAdmin(), UserApi.UserUpdateView)
	r.PUT("/user_password", middleware.JwtAuth(), UserApi.UserUpdatePasswordView)
	r.GET("/user_list", middleware.JwtAuth(), UserApi.UserListView)
	r.POST("/logout", middleware.JwtAuth(), UserApi.UserLogoutView)
	r.DELETE("/user", middleware.JwtAdmin(), UserApi.UserRemoveView)
}
