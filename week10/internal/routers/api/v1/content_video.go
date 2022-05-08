package v1

import (
	"github.com/gin-gonic/gin"
	"parent-api-go/global"
	"parent-api-go/internal/service"
	"parent-api-go/pkg/app"
	"parent-api-go/pkg/context"
	"parent-api-go/pkg/errcode"
	"strconv"
)

type ContentVideo struct{}

func NewContentVideo() ContentVideo {
	return ContentVideo{}
}

// @title    获取分类的视频列表
// @description   获取某个分类下的视频列表
// @auth      zhanglijie01     时间（2021/4/28   20:57 ）
// @param     subject_id       int         "分类编号"
// @Success 200 {object} model.ContentVideoSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/content-video [get]
func (ContentVideo) List(c *gin.Context) {
	param := service.ContentVideoListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	videos, totalRows, err := svc.GetContentVideoList(&param, &pager, "id DESC", 0)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetContentVideoList err:%v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponseList(videos, totalRows)
}

func (ContentVideo) Detail(c *context.AppContext) {
	id := c.Context.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := app.NewResponse(c.Context)

	if idInt <= 0 {
		global.Logger.Error(c, "param[id] invalid", id)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("param[id] invalid " + id))
		return
	}

	svc := service.New(c.Request.Context())
	r, err := svc.GetVideoDetail(int32(idInt))
	if err != nil {
		global.Logger.Errorf(c, "svc.GetContentVideoList err:%v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	response.ToResponse(r)
}
