package routers

import (
	"github.com/gin-gonic/gin"
	"parent-api-go/global"
	"parent-api-go/internal/middleware"
	"parent-api-go/internal/routers/api"
	v1 "parent-api-go/internal/routers/api/v1"
	"parent-api-go/pkg/context"
	"parent-api-go/pkg/limiter"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	//r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	//r.Use(middleware.Translations())
	r.Use(middleware.Cors())

	contentVideo := v1.NewContentVideo()
	home := v1.NewHome()
	balance := v1.NewBalance()
	appIndex := v1.NewAppIndex()
	contentArticle := v1.NewContentArticle()
	invoice := v1.NewInvoice()

	r.GET("/debug/vars", api.Expvar)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.Starter(), middleware.Cors())
	{
		// 获取文章列表
		apiv1.GET("/content_video/list", contentVideo.List)
		apiv1.GET("/content_video/detail/:id", middleware.Authenticate(), middleware.SafeAuthReq(), context.Handle(contentVideo.Detail))
		//用户中心展示的卡片
		apiv1.GET("/app/content_video/list", contentVideo.List)
		//apiv1.GET("/content_video/detail/:id",middleware.Authenticate(),middleware.SafeAuthReq(),context.Handle(contentVideo.Detail))
		apiv1.GET("/app/content_video/detail/:id", context.Handle(contentVideo.Detail))
		apiv1.GET("/app/content_article/list", contentArticle.List)
		apiv1.GET("/app/content_article/detail/:id", context.Handle(contentArticle.Detail))
		apiv1.GET("/app/index", context.Handle(appIndex.Index))
		apiv1.GET("/message/ad/screen", context.Handle(appIndex.GetScreenAd))
		//apiv1.POST("/home/cards", home.Cards)
		// 发票管理
		apiv1.GET("/invoice/order/list", context.Handle(invoice.GetUserInvoiceList))

		//用户中心展示的卡片内容块
		apiv1.POST("/home/cards", middleware.Authenticate(), middleware.SafeAuthReq(), context.Handle(home.Cards))
		//用户中心帐户余额
		apiv1.POST("/home/balances", middleware.Authenticate(), middleware.SafeAuthReq(), context.Handle(home.UserBalance))
		/// 对应原php的/v1/user/balance
		apiv1.POST("/profile/balances", middleware.Authenticate(), middleware.SafeAuthReq(), context.Handle(balance.UserBalance))

		//课时余额替换原php的/v1/app/profile/userbalance
		apiv1.POST("/profile/class/balances", middleware.Authenticate(), middleware.SafeAuthReq(), context.Handle(balance.UserClassBalance))

		//用户够买的课时包列表 ，对应php的/v1/app/profile/userpcpulist
		apiv1.POST("/profile/pcpu/list", middleware.Authenticate(), middleware.SafeAuthReq(), context.Handle(balance.GetStudentPcpuRecords))
		//用户购买的某个课时包的交流流水记录。 对应php的/v1/student/pcpu/balance/detail
		apiv1.POST("/profile/pcpu/transflow/record", middleware.Authenticate(), middleware.SafeAuthReq(), context.Handle(balance.GetStudentPcpuTransactionFlowRecord))

	}
	return r
}
