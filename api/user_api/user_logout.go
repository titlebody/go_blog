package user_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model/res"
	"go_blog/service"
	"go_blog/utils/jwts"
)

func (UserAPI) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")

	//需要计算距离现在的过期时间
	err := service.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("退出登录失败", c)
		return
	}
	res.OkWithMessage("退出登录成功", c)

}
