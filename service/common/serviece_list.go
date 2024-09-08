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
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认时间往前排
		//option.Sort = "created_at asc"  //默认时间往后排
	}
	query := DB.Where(mode)
	//总数
	count = query.Select("id").Find(&list).RowsAffected
	//这里的query会受上面查询的影响，需要复位
	query = DB.Where(mode)
	// 当前第几页
	offset := option.Limit * (option.Page - 1)
	if offset < 0 {
		offset = 0
	}
	if option.Limit <= 0 {
		option.Limit = int(count)
	}
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
