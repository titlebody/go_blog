package user_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
)

func (UserAPI) UserDetailView(c *gin.Context) {
	// 获取用户ID
	id := c.Param("id")
	// 获取用户信息
	var user model.UserModel
	if err := global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	res.OkWithData(user, c)
}
