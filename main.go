package main

import (
	"go_blog/core"
	"go_blog/global"
	"go_blog/router"
)

func main() {
	// 配置文件初始化
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	// 数据库初始化
	global.DB = core.InitGorm()
	// 路由
	r := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("运行在：%s", addr)
	r.Run(addr)

}
