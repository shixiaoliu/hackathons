package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			// 如果没有 Origin 头，使用通配符允许所有源
			origin = "*"
		}

		// 设置 CORS 头
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Header("Access-Control-Max-Age", "86400") // 预检请求结果缓存24小时

		// 处理 OPTIONS 预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 返回 204 No Content
			return
		}

		c.Next()
	}
}
