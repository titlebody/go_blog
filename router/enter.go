package router

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	})
	return r
}
