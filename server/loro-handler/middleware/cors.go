package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetCorsHeader(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Max-Age", "10000")
	c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Authorization,Content-Type,Accept")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
