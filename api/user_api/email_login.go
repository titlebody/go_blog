package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
	"go_blog/utils/jwts"
	"go_blog/utils/pwd"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserAPI) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var userModel model.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 用户不存在
		global.Log.Warn("用户不存在")
		res.FailWithMessage("用户名或者密码错误", c)
		return
	}
	// 密码校验
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	fmt.Println(isCheck)
	if !isCheck {
		res.FailWithMessage("用户名或者密码错误", c)
		return
	}
	token, err := jwts.CreateToken(jwts.JwtPayLoad{
		UserID:   userModel.ID,
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
	})
	if err != nil {
		res.FailWithMessage("token生成失败", c)
		return
	}
	res.OkWithData(gin.H{
		"token": token,
		"name":  userModel.NickName,
	}, c)

}
