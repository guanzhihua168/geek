package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"parent-api-go/global"
	"parent-api-go/internal/service"
	"parent-api-go/pkg/app"
	"parent-api-go/pkg/context"
	"parent-api-go/pkg/errcode"
	"strconv"
)

type ContentArticle struct{}

func NewContentArticle() ContentArticle {
	return ContentArticle{}
}

func (ContentArticle) List(c *gin.Context) {
	param := service.ContentArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	videos, totalRows, err := svc.GetContentArticleList(&param, &pager,
		0, []string{"id", "subject_name", "title", "subtitle", "cover", "created_at"}, "id DESC")
	if err != nil {
		global.Logger.Errorf(c, "svc.GetContentVideoList err:%v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponseList(videos, totalRows)
}

func (ContentArticle) Detail(c *context.AppContext) {
	id := c.Context.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := app.NewResponse(c.Context)
	fmt.Println(id)

	if idInt <= 0 {
		global.Logger.Error(c, "param[id] invalid", id)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("param[id] invalid " + id))
		return
	}

	svc := service.New(c.Request.Context())
	r, err := svc.GetArticleDetail(int32(idInt))
	if err != nil {
		global.Logger.Errorf(c, "svc.List err:%v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponse(r)
}
