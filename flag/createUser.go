package flag

import (
	"fmt"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/c_type"
	"go_blog/utils/pwd"
)

func CreateUser(permissions string) {
	// 创建用户的逻辑
	// 用户名 昵称 密码 确认密码 邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)

	fmt.Printf("请输入用户名:")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称:")
	fmt.Scan(&nickName)
	fmt.Printf("请输入邮箱:")
	fmt.Scan(&email)
	fmt.Printf("请输入密码:")
	fmt.Scan(&password)
	fmt.Printf("请再次输入密码:")
	fmt.Scan(&rePassword)

	// 判断用户名是否存在
	var userModel model.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		global.Log.Error("用户名已存在,请重新输入", err)
		return
	}
	//校验两次密码
	if password != rePassword {
		global.Log.Error("两次密码不一致,请重新输入", err)
		return
	}
	//对密码进行hash
	hashPwd := pwd.HashPwd(password)

	role := c_type.PermissionUser

	if permissions == "admin" {
		role = c_type.PermissionAdmin

	}
	// 头像
	// 1.默认头像
	avatar := "uploads/avatar/default.webp"
	// 2.随机头像
	//入库
	err = global.DB.Create(&model.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: c_type.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户:%s创建成功!", userName)

	fmt.Println(userName, nickName, password, rePassword, email)

}
