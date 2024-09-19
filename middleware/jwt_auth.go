package middleware

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model/c_type"
	"go_blog/model/res"
	"go_blog/service/redis"
	"go_blog/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		// 判断是否在redis中
		ok := redis.CheckLogout(token)
		global.Log.Error(ok)
		if ok {
			res.FailWithMessage("token已过期", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
	}

}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		// 判断是否在redis中
		ok := redis.CheckLogout(token)
		global.Log.Error(ok)
		if ok {
			res.FailWithMessage("token已过期", c)
			c.Abort()
			return
		}
		// 登录的用户
		if claims.Role != int(c_type.PermissionAdmin) {
			res.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}

}
