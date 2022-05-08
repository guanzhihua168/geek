package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"parent-api-go/pkg/redis"
	"parent-api-go/pkg/response"
	"parent-api-go/pkg/util"
	"strconv"
	"time"
)

func SafeAuthReq() gin.HandlerFunc {
	return func(c *gin.Context) {
		ac, e := util.GetAppContext(c)
		if e != nil {
			c.JSON(response.Opps("系统错误"))
			c.Abort()
			return
		}
		if ac.AuthId == 0 {
			c.JSON(response.Opps("您无权操作"))
			c.Abort()
			return
		}

		sign := "[service-activity]:" + c.Request.URL.Path + ":" + strconv.Itoa(int(ac.AuthId))

		l := redis.RedisGo{}.Mutex(sign, 10*time.Second)
		if e := l.Lock(); e != nil {
			logrus.WithField("sign", sign).Info("repeat request" + e.Error())
			c.JSON(response.Opps("系统繁忙，请勿重复请求"))
			c.Abort()
			return
		}
		defer l.Unlock()
		c.Next()
	}
}
