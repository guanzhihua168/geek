package v1

import (
	"parent-api-go/global"
	"parent-api-go/internal/model/mop"
	"parent-api-go/internal/service"
	"parent-api-go/pkg/app"
	"parent-api-go/pkg/context"
	"parent-api-go/pkg/errcode"
	"parent-api-go/pkg/util"
)

type AppIndex struct{}

func NewAppIndex() AppIndex {
	return AppIndex{}
}

func (AppIndex) Index(c *context.AppContext) {
	var (
		err            error
		param          = service.AppIndexRequest{}
		response       = app.NewResponse(c.Context)
		parentLineId   uint8
		banners        []*mop.BannerModel
		videoResults   []*service.VideoResult
		articleResults []*service.ArticleResult
	)

	valid, errs := app.BindAndValid(c.Context, &param)
	if !valid {
		global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	parentLineId, _ = svc.GetParentLineId(c.AuthId)

	banners, err = svc.GetIndexBanners(c.AuthId, param.Platform, param.Device.VApp, parentLineId)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetContentVideoList err:%v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	videoResults, articleResults, err = svc.GetIndexContent(parentLineId)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetIndexContent err:%v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	data := make(map[string]interface{})
	data["banners"] = banners
	data["video_content"] = videoResults
	data["article_content"] = articleResults

	response.ToResponse(data)
}

/**
 * @Description: 获取闪屏广告
 * @receiver AppIndex
 * @param c
 */
func (AppIndex) GetScreenAd(c *context.AppContext) {
	var (
		lineId      uint8
		err         error
		response    = app.NewResponse(c.Context)
		device      = service.RequestDeviceP{}
		pushMessage *mop.PushMessageModel
		platform    int
	)

	valid, errs := app.BindAndValid(c.Context, &device)
	if !valid {
		global.Logger.Error(c.Context, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	home := NewHome()
	flag, err := home.isIOSAuditVersion(c.Context, device)
	if err != nil {
		global.Logger.Error(c, "[home.isIOSAuditVersion] err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	// iOS提审中判断
	if flag {
		response.ToResponse(nil)
		return
	}

	svc := service.New(c.Request.Context())

	// 区分欧美标
	if c.AuthId > 0 {
		lineId, err = svc.GetParentLineId(c.AuthId)
	} else {
		lineId = global.AmeLineId
	}

	studentId := svc.GetStudentId(c.AuthId)
	ua := c.GetHeader("user-agent")
	mobileOS := util.GetMobileOS(ua, device.AKey)

	pushMessage, err = svc.GetAdFirst(studentId, c.AuthId, true, lineId, mobileOS)
	if err != nil {
		global.Logger.Error(c, "[svc.GetAdFirst] err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	if pushMessage == nil || pushMessage == (&mop.PushMessageModel{}) {
		platform = 0
	}

	if pushMessage != nil && pushMessage.Platform != 3 {
		if util.IsIOS(ua, device.AKey) && platform != 1 {
			pushMessage = nil
		} else if util.IsAndroid(ua, device.AKey) && platform != 2 {
			pushMessage = nil
		}
	}

	response.ToResponse(svc.FormatScreenAd(pushMessage))
}

/**
 * @Description: 获取弹屏广告
 * @receiver AppIndex
 * @param c
 */
func (AppIndex) GetAd(c *context.AppContext) {
	var (
		lineId      uint8
		err         error
		response    = app.NewResponse(c.Context)
		device      = service.RequestDeviceP{}
		pushMessage *mop.PushMessageModel
	)

	valid, errs := app.BindAndValid(c.Context, &device)
	if !valid {
		global.Logger.Error(c.Context, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	home := NewHome()
	flag, err := home.isIOSAuditVersion(c.Context, device)
	if err != nil {
		global.Logger.Error(c, "[home.isIOSAuditVersion] err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	// iOS提审中判断
	if flag {
		response.ToResponse(nil)
		return
	}

	svc := service.New(c.Request.Context())

	// 区分欧美标
	if c.AuthId > 0 {
		lineId, err = svc.GetParentLineId(c.AuthId)
	} else {
		lineId = global.AmeLineId
	}

	studentId := svc.GetStudentId(c.AuthId)
	ua := c.GetHeader("user-agent")
	mobileOS := util.GetMobileOS(ua, device.AKey)

	pushMessage, err = svc.GetAdFirst(studentId, c.AuthId, false, lineId, mobileOS)
	if err != nil {
		global.Logger.Error(c, "[svc.GetAdFirst] err: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	response.ToResponse(svc.FormatAd(pushMessage))
}
