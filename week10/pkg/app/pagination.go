package app

import (
	"github.com/gin-gonic/gin"
	"parent-api-go/global"
	"parent-api-go/pkg/convert"
)

//获取请求的页码
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

//获取请求中的分页条数
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.DefaultPageSize {
		return global.AppSetting.DefaultPageSize
	}

	return pageSize
}

//获取分页产生的offset
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
