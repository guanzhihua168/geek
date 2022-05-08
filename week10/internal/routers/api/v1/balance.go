package v1

import (
	"parent-api-go/global"
	"parent-api-go/internal/repository"
	"parent-api-go/internal/service"
	"parent-api-go/pkg/app"
	"parent-api-go/pkg/context"
	"parent-api-go/pkg/errcode"
)

type Balance struct{}

func NewBalance() Balance {
	return Balance{}
}

type UserClassBalanceResult struct {
	RestClass        int `json:"rest_class"`
	RestMainClass    int `json:"rest_main_class"`
	RestNotMainClass int `json:"rest_not_main_class"`
}

/**
 * @Description: 用户个人中心相关用户余额相关信息
 * @receiver p
 * @param c
 */
func (p Balance) UserBalance(c *context.AppContext) {

	svc := service.New(c.Request.Context())
	response := app.NewResponse(c.Context)
	userID := c.AuthId

	parentBalanceInfo, err := svc.GetParentBalance(userID)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorParentBalance)
	}

	response.ToResponse(parentBalanceInfo)
}

func (p Balance) UserClassBalance(c *context.AppContext) {
	svc := service.New(c.Request.Context())
	response := app.NewResponse(c.Context)
	userID := c.AuthId

	var classBalance UserClassBalanceResult

	userRepos := &repository.UserRepos{}
	userRepos.Init()
	studentID := userRepos.GetStuIdByUid(userID)
	lineID, err := userRepos.GetLineId(userID)

	if lineID == global.BUSINESS_LINE_UK {
		classBalance.RestClass, err = svc.GetStudentRestClass(studentID, lineID)
		if err != nil {
			global.Logger.Errorf(c.Context, "svc.GetStudentRestClass errs: %v", err)
			response.ToErrorResponse(errcode.ServerError)
			return
		}
		classBalance.RestMainClass, err = svc.GetStudentRestClassByByMainClass(studentID, lineID, 1)
		if err != nil {
			global.Logger.Errorf(c.Context, "svc.GetStudentRestClassByByMainClass errs: %v", err)
			response.ToErrorResponse(errcode.ServerError)
			return
		}
		classBalance.RestNotMainClass, err = svc.GetStudentRestClassByByMainClass(studentID, lineID, 0)
		if err != nil {
			global.Logger.Errorf(c.Context, "svc.GetStudentRestClassByByMainClass errs: %v", err)
			response.ToErrorResponse(errcode.ServerError)
			return
		}
	} else {
		classBalance.RestClass, err = svc.GetWalletStudentRestClass(studentID)
		if err != nil {
			global.Logger.Errorf(c.Context, "GetWalletStudentRestClass errs: %v", err)
			response.ToErrorResponse(errcode.ServerError)
			return
		}
		classBalance.RestMainClass, err = svc.GetWalletStudentRestClassByMainClass(studentID, 1)
		if err != nil {
			global.Logger.Errorf(c.Context, "svc.GetWalletStudentRestClassByMainClass errs: %v", err)
			response.ToErrorResponse(errcode.ServerError)
			return
		}
		classBalance.RestNotMainClass, err = svc.GetWalletStudentRestClassByMainClass(studentID, 0)
		if err != nil {
			global.Logger.Errorf(c.Context, "svc.GetStudentRestClassByByMainClass errs: %v", err)
			response.ToErrorResponse(errcode.ServerError)
			return
		}
	}

	response.ToResponse(classBalance)
}

/**
 * @Description:获取学生购买的课时包记录
 * @receiver p
 * @param c
 */
func (p Balance) GetStudentPcpuRecords(c *context.AppContext) {
	param := service.PcpuRecordsRequest{}
	response := app.NewResponse(c.Context)
	valid, errs := app.BindAndValid(c.Context, &param)
	if !valid {
		global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	userID := c.AuthId
	userRepos := &repository.UserRepos{}
	userRepos.Init()
	studentID := userRepos.GetStuIdByUid(userID)
	lineID, err := userRepos.GetLineId(userID)
	if err != nil {
		global.Logger.Error(c, "userRepos.GetLineId errs: %v", errs)
		response.ToErrorResponse(errcode.ServerError)
	}

	var records []service.PcpuResult
	if lineID == global.BUSINESS_LINE_US {
		records, err = svc.GetStudentUSPcpuList(studentID, param.IsMainClass)
		if err != nil {
			global.Logger.Errorf(c, "svc.GetStudentUSPcpuList err: %v", err)
			response.ToErrorResponse(errcode.ErrorStudentBalanceRecord)
		}
	}

	response.ToResponse(records)
}

//用户购买的某个pcpu的交易流水的记录
func (p Balance) GetStudentPcpuTransactionFlowRecord(c *context.AppContext) {
	param := service.PcpuFlowDetailRequest{}
	response := app.NewResponse(c.Context)
	valid, errs := app.BindAndValid(c.Context, &param)
	if !valid {
		global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())

	userID := c.AuthId
	userRepos := &repository.UserRepos{}
	userRepos.Init()
	studentID := userRepos.GetStuIdByUid(userID)
	flowRecords, err := svc.GetStudentPcpuTransactionFlowRecord(studentID, param.CourseType, param.Capacity, param.CourseTextbookType)
	if err != nil {
		global.Logger.Error(c, "svc.GetStudentPcpuTransactionFlowRecord errs: %v", errs)
		response.ToErrorResponse(errcode.ServerError)
	}

	response.ToResponse(flowRecords)
}

/**
 * @Description: 家长用户的货币余额交易流水记录
 * @receiver p
 * @param c
 */
func (p Balance) GetParentBalanceTransFlowRecord(c *context.AppContext) {

	param := service.ParentBalanceTransFlowRecordsRequest{}
	response := app.NewResponse(c.Context)
	valid, errs := app.BindAndValid(c.Context, &param)
	if !valid {
		global.Logger.Error(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	userID := c.AuthId
	records, err := svc.GetParentBalanceTransFlowRecord(userID, param.Type)
	if err != nil {
		global.Logger.Error(c, "svc.GetParentBalanceTransFlowRecord errs: %v", errs)
		response.ToErrorResponse(errcode.ServerError)
	}

	response.ToResponse(records)
}
