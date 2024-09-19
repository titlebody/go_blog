package user_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/model"
	"go_blog/model/c_type"
	"go_blog/model/res"
	"go_blog/service/common"
	"go_blog/utils/desens"
	"go_blog/utils/jwts"
)

func (UserAPI) UserListView(c *gin.Context) {
	// 如何判断是超级管理员
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page model.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(model.UserModel{}, common.Option{
		PageInfo: page,
	})
	var users []model.UserModel
	for _, user := range list {
		if claims.Role != int(c_type.PermissionAdmin) {
			//非管理员
			user.UserName = ""

		}
		// 手机号脱敏
		//176****2311
		user.Tel = desens.Desensitization(user.Tel)
		// 邮箱脱敏
		// 2****@qq.com
		user.Email = desens.DesensitizationEamil("2387360024@qq.com")

		users = append(users, user)
	}

	res.OkWithList(users, count, c)
}
