package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"parent-api-go/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	//页码
	Page int `json:"page"`
	//每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}
type ResDataList struct {
	List  interface{} `json:"list"`
	Pager Pager       `json:"pager"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

//生成响应数据
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	response := gin.H{"code": http.StatusOK, "message": "success", "data": data}
	r.Ctx.JSON(http.StatusOK, response)
}

// 带分页输出内容返回
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"total":   totalRows,
		"data": ResDataList{
			List: list,
			Pager: Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	})
}

//返回带有错误的响应数据
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "message": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
