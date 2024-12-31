package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求来源
		origin := c.Request.Header.Get("Origin")

		// 检查是否允许的来源，这里需要你自己定义一个允许的来源列表
		allowedOrigins := []string{"http://localhost:8080", "https://blog.huchenxi.fun", "https://www.huchenxi.fun"}
		isAllowedOrigin := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				isAllowedOrigin = true
				break
			}
		}

		if isAllowedOrigin {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
