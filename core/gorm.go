package core

import (
	"fmt"
	"go_blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	//判断配置文件
	if global.Config.Mysql.Host == "" {
		log.Panicln("未配置mysql,取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	//判断配置环境
	if global.Config.System.Env == "debug" {
		// 开发环境显示所有sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) //只打印错误sql
	}
	// 连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		fmt.Println("连接失败", dsn)
		log.Panicln(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间
	return db
}
