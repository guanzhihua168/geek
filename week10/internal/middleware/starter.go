package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"parent-api-go/global"
	"parent-api-go/pkg/context"
)

func Starter() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logrus.WithField("trace_id", "wait")
		ac := &context.AppContext{
			c,
			log,
			0,
		}
		c.Set(global.AppContext, ac)
		c.Next()
	}
}
