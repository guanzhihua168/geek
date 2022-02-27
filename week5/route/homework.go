package route

import (
	v1 "geek/week5/api/v1"
	"geek/week5/middleware"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.Use(middleware.LimitRate)

	{
		r.GET("/homework/index", v1.List)
	}

}
