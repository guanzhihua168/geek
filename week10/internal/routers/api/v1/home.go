package v1

import (
	"github.com/gin-gonic/gin"
	"parent-api-go/global"
	"parent-api-go/internal/repository"
	"parent-api-go/internal/service"
	"parent-api-go/pkg/app"
	"parent-api-go/pkg/context"
	"parent-api-go/pkg/errcode"
	"parent-api-go/pkg/setting"
	"strings"
)

// 提审中
const IsAuditing int = 1

type Home struct{}

func NewHome() Home {
	return Home{}
}

type UserBalanceInfo struct {
	UserBalance        int    `json:"user_balance"`
	CoinBalance        int    `json:"coin_balance"`
	DisplayCoinBalance string `json:"display_coin_balance"`
	RestClass          int    `json:"rest_class"`
	RestMainClass      int    `json:"rest_main_class"`
	RestNotMainClass   int    `json:"rest_not_main_class"`
}

// 用户个人中心相关用户余额相关信息
func (H Home) UserBalance(c *context.AppContext) {
	var err error
	svc := service.New(c.Request.Context())
	response := app.NewResponse(c.Context)
	userID := c.AuthId
	var userBalanceInfo UserBalanceInfo
	userBalanceInfo.UserBalance, err = svc.GetUserMoneyBalance(userID)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	userBalanceInfo.CoinBalance, err = svc.GetUserCoinBalance(userID)
	if err != nil {
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	userBalanceInfo.DisplayCoinBalance = svc.GetDisplayCoinBalance(userBalanceInfo.CoinBalance)
	userRepos := &repository.UserRepos{}
	userRepos.Init()
	studentID := userRepos.GetStuIdByUid(userID)
	lineID, err := userRepos.GetLineId(userID)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetUserLineID)
		return
	}
	if lineID == global.BUSINESS_LINE_UK {
		userBalanceInfo.RestClass, err = svc.GetStudentRestClass(studentID, lineID)
		if err != nil {
			response.ToErrorResponse(errcode.ServerError)
			return
		}
		userBalanceInfo.RestMainClass, err = svc.GetStudentRestClassByByMainClass(studentID, lineID, 1)
		if err != nil {
			response.ToErrorResponse(errcode.ServerError)
			return
		}
		userBalanceInfo.RestNotMainClass, err = svc.GetStudentRestClassByByMainClass(studentID, lineID, 0)
		if err != nil {
			response.ToErrorResponse(errcode.ServerError)
			return
		}
	} else {
		userBalanceInfo.RestClass, err = svc.GetWalletStudentRestClass(studentID)
		if err != nil {
			global.Logger.Errorf(c.Context, "GetWalletStudentRestClass errs: %v", err)
			response.ToErrorResponse(errcode.ServerError)
			return
		}
		userBalanceInfo.RestMainClass, err = svc.GetWalletStudentRestClassByMainClass(studentID, 1)
		if err != nil {
			global.Logger.Errorf(c.Context, "svc.GetWalletStudentRestClassByMainClass errs: %v", err)
			response.ToErrorResponse(errcode.ServerError)
			return
		}
		userBalanceInfo.RestNotMainClass, err = svc.GetWalletStudentRestClassByMainClass(studentID, 0)
		if err != nil {
			global.Logger.Errorf(c.Context, "svc.GetStudentRestClassByByMainClass errs: %v", err)
			response.ToErrorResponse(errcode.ServerError)
			return
		}
	}

	response.ToResponse(userBalanceInfo)
}

//用户个人中心 展示的相关卡片
func (H Home) Cards(c *context.AppContext) {
	param := service.HomeCardsRequest{}
	response := app.NewResponse(c.Context)
	valid, errs := app.BindAndValid(c.Context, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	userID := c.AuthId
	userRepos := &repository.UserRepos{}
	userRepos.Init()
	parentLineID, err := userRepos.GetLineId(userID)
	if err != nil {
		global.Logger.Errorf(c, "userRepos.GetLineId errs: %v", errs)
		response.ToErrorResponse(errcode.ServerError)
	}

	s, err := H.getCardsSettingByLineID(parentLineID)
	if err != nil {
		global.Logger.Errorf(c, "H.getCardsSettingByLineID errs: %v", errs)
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	reqAppVersion := param.Device.VApp
	var AccountCardsSetting *setting.AccountCardsSettingS
	if reqAppVersion != "" {
		err = s.ReadSection(reqAppVersion, &AccountCardsSetting)
		if err != nil {
			response.ToErrorResponse(errcode.ServerError)
			return
		}
	}
	if AccountCardsSetting == nil {
		err = s.ReadSection("default", &AccountCardsSetting)
	}
	//ios 提审时过滤
	AccountCardsSetting.Balance.Items, _ = H.filterAuditHiddenBalanceItems(AccountCardsSetting.Balance.Items, AccountCardsSetting.IosAuditHidden)
	AccountCardsSetting.CardList, _ = H.filterAuditHiddenCardListItems(AccountCardsSetting.CardList, AccountCardsSetting.IosAuditHidden)
	response.ToResponse(AccountCardsSetting)
}

/**
 * @Description:获取不同业务线的卡片配置
 * @receiver H
 * @param lineID
 * @return *setting.Setting
 * @return error
 */
func (H Home) getCardsSettingByLineID(lineID int) (*setting.Setting, error) {
	var s *setting.Setting
	var err error
	if lineID == global.BUSINESS_LINE_UK {
		s, err = setting.NewUKAccountCardsSetting()
		if err != nil {
			return nil, err
		}
	} else {
		s, err = setting.NewUSAccountCardsSetting()
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

//根据提ios提审核配置，过滤掉不显示用户帐帐户余额相关的卡片
func (H Home) filterAuditHiddenBalanceItems(balanceItems []setting.BalanceItemStruct, auditHidden []string) ([]setting.BalanceItemStruct, error) {
	var tempBalanceItems []setting.BalanceItemStruct
	for i := 0; i < len(balanceItems); i++ {
		if !inArray(balanceItems[i].ItemCode, auditHidden) {
			tempBalanceItems = append(tempBalanceItems, balanceItems[i])
		}
	}
	return tempBalanceItems, nil
}

func (H Home) filterAuditHiddenCardListItems(cardList []setting.CardListStruct, auditHidden []string) (
	[]setting.CardListStruct, error) {

	var newCardList []setting.CardListStruct
	for i := 0; i < len(cardList); i++ {
		var newItems []setting.CardListItem
		for j := 0; j < len(cardList[i].Items); j++ {
			if !inArray(cardList[i].Items[j].ItemCode, auditHidden) {
				newItems = append(newItems, cardList[i].Items[j])
			}
		}
		cardList[i].Items = newItems
		if newItems != nil {
			newCardList = append(newCardList, cardList[i])
		}
	}
	return newCardList, nil
}

func inArray(item string, items []string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func (H Home) isIOSAuditVersion(c *gin.Context, device service.RequestDeviceP) (bool, error) {
	//只有ios才有提审
	platform := 1
	svc := service.New(c.Request.Context())
	if strings.ToLower(device.Lang) == "ios" {
		appVersion := device.VApp
		appVersionConfig, err := svc.GetLatestVersionConfig(platform, appVersion)
		if err != nil {
			return false, err
		}
		if appVersionConfig.IsAudition == IsAuditing {
			return true, nil
		}
	}
	return false, nil
}
