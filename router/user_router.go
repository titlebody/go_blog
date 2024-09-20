package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"go_blog/api"
	"go_blog/middleware"
)

var store = cookie.NewStore([]byte("chenxi214dassd"))

func (r RouterGroup) UserRouter() {
	UserApi := api.ApiGroupApp.UserAPI
	r.Use(sessions.Sessions("sessionid", store))
	r.POST("/email_login", UserApi.EmailLoginView)
	r.PUT("/user_role", middleware.JwtAdmin(), UserApi.UserUpdateView)
	r.PUT("/user_password", middleware.JwtAuth(), UserApi.UserUpdatePasswordView)
	r.GET("/user_list", middleware.JwtAuth(), UserApi.UserListView)
	r.POST("/logout", middleware.JwtAuth(), UserApi.UserLogoutView)
	r.POST("/user_bind_email", middleware.JwtAuth(), UserApi.UserBindEmailView)
	r.DELETE("/user", middleware.JwtAdmin(), UserApi.UserRemoveView)
}
