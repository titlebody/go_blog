package core

import (
	"go_blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitGorm() *gorm.DB {
	//判断配置文件
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql,取消gorm连接")
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

	global.MysqlLog = logger.Default.LogMode(logger.Info)

	// 连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表名不加复数
		},
	})
	if err != nil {
		global.Log.Fatalf("[%s]连接失败", dsn)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间
	return db
}
