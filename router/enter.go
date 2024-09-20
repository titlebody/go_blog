package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_blog/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	//注册swagger api相关路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("api")
	rouerGroup := RouterGroup{api}
	rouerGroup.SettingsRouter()
	rouerGroup.ImagesRouter()
	rouerGroup.AdvertRouter()
	rouerGroup.MenuRouter()
	rouerGroup.UserRouter()
	rouerGroup.TagRouter()
	rouerGroup.MessageRouter()
	//系统配置api

	return r
}
