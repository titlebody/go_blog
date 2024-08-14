package main

import (
	"go_blog/core"
	"go_blog/global"
)

func main() {
	// 配置文件初始化
	core.InitConfig()
	// 数据库初始化
	global.DB = core.InitGorm()

}
