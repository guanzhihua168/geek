package middleware

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"parent-api-go/internal/repository"
	"parent-api-go/pkg/response"
	"parent-api-go/pkg/util"

	"strings"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(auth) < 2 {
			c.JSON(response.Opps("Authentication empty"))
			c.Abort()
			return
		}
		baseAuth := util.SafeBase64Replace(auth[1])
		b, e := base64.RawStdEncoding.DecodeString(baseAuth)
		if e != nil {
			c.JSON(response.Opps("Authentication failed"))
			c.Abort()
			return
		}
		ac, e := util.GetAppContext(c)
		if e != nil {
			c.JSON(response.Opps("系统错误"))
			c.Abort()
			return
		}

		repo := repository.UserRepos{}
		ac.CtxInto(&repo)
		if uid := repo.ParseAuthToken(string(b)); uid == 0 {
			c.JSON(response.Opps("Authentication failed"))
			c.Abort()
			return
		} else {
			ac.AuthId = uid
			ac.Log = ac.Log.WithField("uid", uid)
			c.Set("AuthId", uid)
		}
		c.Next()
	}
}
