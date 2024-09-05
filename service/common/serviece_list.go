package common

import (
	"go_blog/global"
	"go_blog/model"
	"gorm.io/gorm"
)

type Option struct {
	model.PageInfo
	Debug bool
}

func ComList[T any](mode T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	//总数
	count = DB.Select("id").Find(&list).RowsAffected
	// 当前第几页
	offset := option.Limit * (option.Page - 1)
	if offset < 0 {
		offset = 0
	}
	if option.Limit == 0 {
		option.Limit = int(count)
	}
	err = DB.Limit(option.Limit).Offset(offset).Find(&list).Error
	return list, count, err
}
