package user_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
	"go_blog/utils/jwts"
	"go_blog/utils/pwd"
	"strconv"
)

type UpdatePasswordRequest struct {
	OldPwd uint `json:"old_pwd" binding:"required" msg:"请输入旧密码"`
	NewPwd uint `json:"new_pwd" binding:"required" msg:"请输入新密码"`
}

func (UserAPI) UserUpdatePasswordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var user model.UserModel

	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	if !pwd.CheckPwd(user.Password, strconv.Itoa(int(cr.OldPwd))) {
		res.FailWithMessage("密码错误", c)
		return
	}

	hashPwd := pwd.HashPwd(strconv.Itoa(int(cr.NewPwd)))
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)
}
