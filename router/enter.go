package router

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	api := r.Group("api")
	rouerGroup := RouterGroup{api}
	rouerGroup.SettingsRouter()
	rouerGroup.ImagesRouter()
	//系统配置api

	return r
}
