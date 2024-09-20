package user_ser

import (
	"errors"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/c_type"
	"go_blog/utils/pwd"
)

func (UserService) CreateUser(userName, nickName, password string, role c_type.Role) error {
	// 判断用户名是否存在
	var userModel model.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		global.Log.Error("用户名已存在,请重新输入", err)
		return errors.New("用户名已存在,请重新输入")
	}
	//对密码进行hash
	hashPwd := pwd.HashPwd(password)

	// 头像
	// 1.默认头像
	avatar := "uploads/avatar/default.webp"
	// 2.随机头像
	//入库
	err = global.DB.Create(&model.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: c_type.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Infof("用户:%s创建成功!", userName)
	return nil
}
