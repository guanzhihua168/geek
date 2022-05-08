package service

import (
	"parent-api-go/global"
	"parent-api-go/internal/repository"
	"parent-api-go/pkg/util"
	"sort"
	"strconv"
	"strings"
)

type PrepaidAccount struct {
	Balance float64 `json:"balance"`
}
type GiftAccount struct {
	Balance float64 `json:"balance"`
}
type ParentBalanceInfo struct {
	PrepaidAccount PrepaidAccount `json:"prepaid_account"`
	GiftAccount    GiftAccount    `json:"gift_account"`
}

type PcpuRecordsRequest struct {
	IsMainClass int            `json:"is_main_class" binding:"required"`
	Device      RequestDeviceP `json:"device" binding:"required"`
}

//学生的购买的某个pcpu的交交易记录
type PcpuFlowDetailRequest struct {
	CourseType         int `json:"course_type"`
	Capacity           int `json:"capacity"`
	CourseTextbookType int `json:"course_textbook_type"`
	Type               int `json:"type"`
}

//家长用户的货币余额交易流水记录
type ParentBalanceTransFlowRecordsRequest struct {
	Type int `json:"type"`
}

/**
 * @Description: 用户购买的课时包信息
 */
type PcpuInfo struct {
	ID                     int    `json:"id"`
	No                     string `json:"no"`
	Type                   int    `json:"type"`
	Capacity               int    `json:"capacity"`
	CourseType             int    `json:"course_type"`
	RestSessionTimes       int    `json:"rest_session_times"`
	CourseTypeText         string `json:"course_type_text"`
	Duration               int    `json:"duration"`
	DurationText           string `json:"duration_text"`
	CourseTextbookType     int    `json:"course_textbook_type"`
	CourseTextbookTypeText string `json:"course_textbook_type_text"`
	TypeFmt                string `json:"type_fmt"`
}

//学生购买的pcpu返回结果结构
type PcpuResult struct {
	Type                   int    `json:"type"`
	CourseType             int    `json:"course_type"`
	CourseTypeText         string `json:"course_type_text"`
	Capacity               int    `json:"capacity"`
	RestSessionTimes       int    `json:"rest_session_times"`
	Duration               int    `json:"duration"`
	DurationText           string `json:"duration_text"`
	CourseTextbookType     int    `json:"course_textbook_type"`
	CourseTextbookTypeText string `json:"course_textbook_type_text"`
}

/**
 * @Description:学生购买的 pcpu变动流水返回结果结构
 */
type PcpuRecordResult struct {
	CreatedAt    int64  `json:"created_at"`
	SessionTimes int    `json:"session_times"`
	TypeText     string `json:"type_text"`
	IoType       int    `json:"io_type"`
}

/**
 * @Description:
 * @receiver svc
 * @param userID
 * @return ParentBalanceInfo
 * @return error
 */
func (svc *Service) GetParentBalance(userID uint32) (ParentBalanceInfo, error) {
	var UserBalanceInfo ParentBalanceInfo
	repo := repository.WalletRepos{}
	balance := repo.GetParentBalance(userID)
	UserBalanceInfo.PrepaidAccount.Balance = balance.ChargeBalanceTotal
	UserBalanceInfo.GiftAccount.Balance = balance.ScholarshipBalanceTotal

	return UserBalanceInfo, nil
}

type UserBalanceRecordResponse struct {
	ID            int32  `json:"id"`
	UserID        uint32 `json:"user_id"`
	Fee           int    `json:"fee"`
	Balance       int    `json:"balance"`
	Type          int    `json:"type"`
	TradeTypeText string `json:"trade_type_text"`
	CreatedAt     string `json:"created_at"`
}

/**
 * @Description:
 * @receiver svc
 * @param studentID
 * @return map[string]PcpuInfo
 * @return error
 */
func (svc *Service) GetStudentPcpus(studentID uint32) (map[string]PcpuInfo, error) {
	repo := repository.WalletRepos{}
	pcpus, err := repo.GetStudentPcpus(studentID)
	if err != nil {
		return nil, err
	}
	mapPcpuInfo := make(map[string]PcpuInfo)

	for _, pcpu := range pcpus {

		if !util.InArrayInts(pcpu.Type, []int{global.PcpuTypeGeneralClass, global.PcpuTypeAllPower}) {
			continue
		}

		newPcpuInfo := PcpuInfo{
			Type:               pcpu.Type,
			CourseType:         pcpu.CourseType,
			Capacity:           pcpu.Capacity,
			Duration:           pcpu.Duration,
			CourseTextbookType: pcpu.CourseTextbookType,
			RestSessionTimes:   pcpu.RestSessionTimes,
		}

		key := strings.Join([]string{
			strconv.Itoa(pcpu.Type),
			strconv.Itoa(pcpu.CourseType),
			strconv.Itoa(pcpu.Capacity),
			strconv.Itoa(pcpu.Duration),
			strconv.Itoa(pcpu.CourseTextbookType)},
			"_")

		if _, ok := mapPcpuInfo[key]; ok {
			// 存在, 对剩余课时累加
			tmpRecord := mapPcpuInfo[key]
			tmpRecord.RestSessionTimes += pcpu.RestSessionTimes
			mapPcpuInfo[key] = tmpRecord
		} else {
			mapPcpuInfo[key] = newPcpuInfo
		}

	}

	return mapPcpuInfo, nil
}

/**
 * @Description: 获取学生美标pcpu 记录
 * @receiver svc
 * @param studentID
 * @param isMainClass
 * @return []PcpuResult
 * @return error
 */
func (svc *Service) GetStudentUSPcpuList(studentID uint32, isMainClass int) ([]PcpuResult, error) {

	pcpuList, err := svc.GetStudentPcpus(studentID)
	if err != nil {
		return nil, err
	}

	var pcpuRecords []PcpuResult

	for _, record := range pcpuList {
		if isMainClass == 1 && !util.InArrayInts(record.CourseType, global.MainClassCourseTypes) {
			continue
		}

		if isMainClass != 1 && util.InArrayInts(record.CourseType, global.MainClassCourseTypes) {
			continue
		}
		//@todo 改为接口取课程名
		courseTypeText := "通用课时"
		if record.Type == global.PcpuTypeGeneralClass {
			if _, ok := global.CourseTypeNameDict[record.CourseType]; ok {
				// 存在
				courseTypeText = global.CourseTypeNameDict[record.CourseType]
			} else {
				courseTypeText = ""
			}
		}

		if record.TypeFmt != "" {
			record.CourseTypeText = record.TypeFmt
		} else {
			record.CourseTypeText = courseTypeText
		}
		record.DurationText = ""
		if _, ok := global.CPUDuration[record.Duration]; ok {
			// 存在
			record.DurationText = global.CPUDuration[record.Duration]
		}

		record.CourseTextbookTypeText = ""
		if _, ok := global.CourseTextBookType[record.CourseTextbookType]; ok {
			record.CourseTextbookTypeText = global.CourseTextBookType[record.CourseTextbookType]
		}
		pcpuRecordResult := PcpuResult{
			Type:                   record.Type,
			CourseType:             record.CourseType,
			CourseTypeText:         record.CourseTypeText,
			Capacity:               record.Capacity,
			RestSessionTimes:       record.RestSessionTimes,
			Duration:               record.Duration,
			DurationText:           record.DurationText,
			CourseTextbookType:     record.CourseTextbookType,
			CourseTextbookTypeText: record.CourseTextbookTypeText,
		}

		pcpuRecords = append(pcpuRecords, pcpuRecordResult)
	}

	return pcpuRecords, nil
}

/**
 * @Description: 获取学生pcpu 记录详情
 * @receiver svc
 * @param studentID
 * @param courseType
 * @param capacity
 * @param courseTextbookType
 * @return []PcpuRecordResult
 * @return error
 */
func (svc *Service) GetStudentPcpuTransactionFlowRecord(studentID uint32, courseType, capacity, courseTextbookType int) ([]PcpuRecordResult, error) {
	repo := repository.WalletRepos{}
	pcpus, err := repo.GetStudentPcpus(studentID)
	if err != nil {
		return nil, err
	}
	records, err := repo.GetStudentPcpuRecords(studentID)
	if err != nil {
		return nil, err
	}

	var ids []int
	//找出符合条件的PcpuID
	for _, pcpu := range pcpus {
		if !util.InArrayInts(pcpu.Type, []int{global.PcpuTypeGeneralClass, global.PcpuTypeAllPower}) {
			continue
		}
		if pcpu.CourseType != courseType {
			continue
		}
		if pcpu.Capacity != capacity {
			continue
		}
		if pcpu.CourseTextbookType != courseTextbookType {
			continue
		}
		ids = append(ids, pcpu.ID)
	}

	//过滤,只符合条件的pcpu record
	var tmpRecords []repository.PcpuRecord

	for _, record := range records {
		if !util.InArrayInts(record.PcpuID, ids) {
			continue
		}
		tmpRecords = append(tmpRecords, record)
	}
	//sort.SliceStable(tmpRecord, func(i, j int) bool { return tmpRecord[i].CreatedAt > tmpRecord[i].CreatedAt }) // 按年龄降序排序
	var recordsResult []PcpuRecordResult
	for _, record := range tmpRecords {
		var ioType int
		if util.InArrayInts(record.Type, global.PcpuBalancePlusTypeList) {
			ioType = global.PcpuBalanceIoPlus
		} else {
			ioType = global.PcpuBalanceIoMinus
		}

		//时间转为微秒格式
		createdAtMicro := record.CreatedAt.Unix() * 1000

		result := PcpuRecordResult{CreatedAt: createdAtMicro, SessionTimes: record.SessionTimes, TypeText: record.TypeFmt, IoType: ioType}
		recordsResult = append(recordsResult, result)
	}

	// 按创建时间降序排序
	sort.Slice(recordsResult, func(i, j int) bool { return recordsResult[i].CreatedAt > recordsResult[j].CreatedAt })

	return recordsResult, nil
}

/**
 * @Description: 家长用户的货币余额交易流水记录
 * @receiver svc
 * @param userID
 * @param recordType
 * @return []*UserBalanceRecordResponse
 * @return error
 */
func (svc *Service) GetParentBalanceTransFlowRecord(userID uint32, recordType int) ([]*UserBalanceRecordResponse, error) {
	records, err := svc.daoOldBossSlave.GetParentBalanceTransFlowRecords(userID, recordType)
	if err != nil {
		return nil, err
	}
	var recordsResponse []*UserBalanceRecordResponse
	for _, record := range records {
		var tradeTypeText string
		if _, ok := global.UserBalanceTradeType[record.TradeType]; ok {
			tradeTypeText = global.UserBalanceTradeType[record.TradeType]
		}

		recordRes := UserBalanceRecordResponse{
			ID:            record.ID,
			Type:          record.Type,
			Balance:       record.Balance / 100,
			Fee:           record.Fee / 100,
			TradeTypeText: tradeTypeText,
		}
		recordsResponse = append(recordsResponse, &recordRes)
	}
	return recordsResponse, nil
}
