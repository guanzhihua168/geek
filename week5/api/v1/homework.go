package v1

import "github.com/gin-gonic/gin"

func List(c *gin.Context) {
	c.JSON(200, gin.H{"id": 1, "week": 5, "name": "gzh", "content": "limit request"})
}
