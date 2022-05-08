package response

import "github.com/gin-gonic/gin"

func Opps(msg string, c ...int) (int, gin.H) {
	code := 3000
	if len(c) > 0 {
		code = c[0]
	}
	return 200, gin.H{
		"code":    code,
		"message": msg,
	}
}

func Okay(data ...interface{}) (int, gin.H) {
	c := 200
	if len(data) == 1 {
		return 200, gin.H{
			"code": c,
			"data": data[0],
		}
	} else if len(data) >= 1 {
		return 200, gin.H{
			"code": data[1],
			"data": data[0],
		}
	} else {
		return 200, gin.H{
			"code": c,
			"data": make(map[string]interface{}),
		}
	}
}
