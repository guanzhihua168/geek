package service

import (
	"parent-api-go/global"
	"parent-api-go/internal/repository"
	"parent-api-go/pkg/util"
	"strconv"
)

type HomeCardsRequest struct {
	ID     string         `json:"id" binding:"required"`
	Device RequestDeviceP `json:"device" binding:"required"`
}

func (svc *Service) GetUserMoneyBalance(userID uint32) (int, error) {
	moneyBalance, err := svc.daoOldBossSlave.GetTotalMoneyBalanceByUserID(userID)
	if err != nil {
		return 0, err
	}
	return moneyBalance, nil
}

func (svc *Service) GetUserCoinBalance(userID uint32) (int, error) {
	repo := repository.ShopRepos{}
	coinBalance := repo.GetUserCoinBalance(userID)

	return coinBalance, nil
}

func (svc *Service) GetDisplayCoinBalance(balance int) string {
	if balance <= 0 {
		return "0"
	} else if balance < global.ShopBaseDisplayCoin {
		return strconv.Itoa(balance)
	} else if balance == global.ShopBaseDisplayCoin {
		return "10M"
	} else {
		return "10M+"
	}
}

/**
 * @Description: 获取学生剩余的总课时数 【欧标使用】
 * @receiver svc
 * @param studentID
 * @param businessLineID
 * @return int
 * @return error
 */
func (svc *Service) GetStudentRestClass(studentID uint32, businessLineID int) (int, error) {
	repo := repository.AtomicUserCpuRepos{}
	repo.Init()
	userCpuTotal, err := repo.GetStudentUserCpuTotal(studentID, businessLineID, 1, nil, nil)
	if err != nil {
		return 0, err
	}

	if len(userCpuTotal) > 0 {
		return userCpuTotal[0].RestAmount, nil
	}

	return 0, nil
}

/**
 * @Description: 获取用户主课/副课的剩余课时 [欧标使用]
 * @receiver svc
 * @param studentID
 * @param lineID
 * @param isMainClass
 * @return int
 * @return error
 */
func (svc *Service) GetStudentRestClassByByMainClass(studentID uint32, lineID int, isMainClass int) (int, error) {

	repo := repository.AtomicUserCpuRepos{}
	repo.Init()
	var courseTypes []int
	if isMainClass == 1 {
		courseTypes = global.MainClassCourseTypes
	} else {
		courseTypes = global.NotMainClassCourseTypes
	}

	conditions := &repository.UserCpuTotalConditions{
		CourseType: 1, //需要通过CourseType条件筛选总剩余课时
	}

	userCpuTotal, err := repo.GetStudentUserCpuTotal(studentID, lineID, 1, courseTypes, conditions)
	if err != nil {
		return 0, nil
	}
	if len(userCpuTotal) > 0 {
		return userCpuTotal[0].RestAmount, nil
	}

	return 0, nil
}

/**
 * @Description: 通过Wallet服务学生pcpu记录，计算得到学生剩余的总课时数
 * @receiver svc
 * @param studentID
 * @param businessLineID
 * @return int
 * @return error
 */
func (svc *Service) GetWalletStudentRestClass(studentID uint32) (int, error) {
	studentPcpus, err := svc.GetStudentPcpus(studentID)
	if err != nil {
		return 0, err
	}
	var restSessionTimes int
	for _, pcpu := range studentPcpus {
		restSessionTimes += pcpu.RestSessionTimes
	}
	return restSessionTimes, nil
}

/**
 * @Description: 通过Wallet服务学生pcpu记录，计算得到学生主课/副课剩余的总课时数
 * @receiver svc
 * @param studentID
 * @param isMainClass
 * @return int
 * @return error
 */
func (svc *Service) GetWalletStudentRestClassByMainClass(studentID uint32, isMainClass int) (int, error) {
	studentPcpus, err := svc.GetStudentPcpus(studentID)
	if err != nil {
		return 0, err
	}
	var restSessionTimes int
	for _, pcpu := range studentPcpus {
		if isMainClass == 1 && util.InArrayInts(pcpu.CourseType, global.MainClassCourseTypes) {
			restSessionTimes += pcpu.RestSessionTimes
		} else if isMainClass != 1 && !util.InArrayInts(pcpu.CourseType, global.MainClassCourseTypes) {
			restSessionTimes += pcpu.RestSessionTimes
		}
	}

	return restSessionTimes, nil
}
