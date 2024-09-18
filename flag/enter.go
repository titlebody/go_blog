package flag

import (
	sys_flag "flag"
)

type Option struct {
	DB   bool
	User string // -u admin  -u user
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")

	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	//sys_flag.Usage()
}
