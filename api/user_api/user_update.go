package user_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/c_type"
	"go_blog/model/res"
)

type UserRole struct {
	Role     c_type.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	UserID   uint        `json:"user_id" binding:"required" msg:"用户ID错误"`
	NickName string      `json:"nick_name"` //防止用户名称非法
}

// UserUpdateView 修改用户权限
func (UserAPI) UserUpdateView(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user model.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改失败", c)
		return
	}
	res.OkWithMessage("修改成功", c)

}
