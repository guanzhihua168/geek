package middleware

import (
	"geek/week5/pkg"
	"github.com/gin-gonic/gin"
)

func LimitRate(c *gin.Context) {
	ip := c.ClientIP()
	limitKey := "limitKey:" + ip
	if !pkg.NewLimiting().LimitReq(limitKey) {
		c.AbortWithStatusJSON(503, gin.H{"message": "error Current IP frequently visited"})
		return
	}

	c.Next()
}
