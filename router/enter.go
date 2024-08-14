package router

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
)

type routerGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	api := r.Group("api")
	rouerGroup := routerGroup{api}
	{
		rouerGroup.SettingsRouter()
	}
	//系统配置api

	return r
}
