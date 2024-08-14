package main

import (
	"fmt"
	"go_blog/config"
)

func main() {
	// 配置文件初始化
	config.InitConfig()

	fmt.Println(config.Config.System)
	fmt.Println(config.Config.Logger)
	fmt.Println(config.Config.Mysql)
	fmt.Println("Hello World")
}
