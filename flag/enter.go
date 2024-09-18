package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
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

func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch v.(type) {
		case string:
			if v.(string) != "" {
				f = true
			}
		case bool:
			if v.(bool) == true {
				f = true
			}
		}
	}
	return f
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
