package user_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/model/c_type"
	"go_blog/model/res"
	"go_blog/service"
)

type UserCreateRequest struct {
	NickName string      `json:"nick_name" binding:"required" msg:"请输入昵称"`  //名称
	UserName string      `json:"user_name" binding:"required" msg:"请输入用户名"` //用户名
	Password string      `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	Role     c_type.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限 1.管理员 2.用户 3.游客
}

func (UserAPI) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	userSer := service.ServiceApp.UserService
	err = userSer.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	res.OkWithMessage("创建成功", c)

}
