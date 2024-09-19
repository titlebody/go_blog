package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"go_blog/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
	Redis    *redis.Client
)
