package global

import (
	"github.com/sirupsen/logrus"
	"go_blog/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// 图片上传的白名单
	WhiteImageList = []string{
		"png",
		"jpg",
		"jpeg",
		"gif",
		"webp",
		"bmp",
		"svg",
		"ico",
		"tiff",
		"tif",
		"jfif",
		"jpe",
		"jif",
	}
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
)
