package main

import (
	"github.com/sirupsen/logrus"
	"go_blog/core"
	"go_blog/global"
)

func main() {
	// 配置文件初始化
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("lalla")
	global.Log.Error("aaa")
	global.Log.Info("ddd")
	logrus.Warnln("lalla")
	logrus.Info("lalla")

	// 数据库初始化
	global.DB = core.InitGorm()

}
