package repository

import (
	"encoding/json"
	"fmt"
	"parent-api-go/global"
	"parent-api-go/pkg/linkerd"
)

type AtomicUserCpuRepos struct {
	repository
}

func (c *AtomicUserCpuRepos) Init() {
	options := linkerdOptions(c.repository)
	options = append(options, linkerd.WithJson())
	c.Remote = linkerd.NewLinkerd(
		global.LinkerdSetting.Host,
		global.LinkerdSetting.AppName,
		global.LinkerdSetting.Token,
		"atomic_user_cpu",
		options...,
	)
}

type AtomicUserCpu struct {
	Duration               int    `json:"duration"`
	CourseTypeText         string `json:"course_type_text"`
	RestAmount             int    `json:"rest_amount"`
	CourseTextbookType     int    `json:"course_textbook_type"`
	DurationText           string `json:"duration_text"`
	Type                   int    `json:"type"`
	Capacity               int    `json:"capacity"`
	CourseTextbookTypeText string `json:"course_textbook_type_text"`
}

type UserCpuTotalConditions struct {
	CourseType int `json:"course_type"`
}
type UserCpuTotalBodyStruct struct {
	StudentID  uint32                  `json:"student_id"`
	LineID     int                     `json:"line_id"`
	Type       int                     `json:"type"`
	CourseType []int                   `json:"course_type"`
	Conditions *UserCpuTotalConditions `json:"conditions"`
	PerPage    int                     `json:"per_page"`
}

//https://yapi.rouchi.com/project/375/interface/api/11408
func (c *AtomicUserCpuRepos) GetStudentUserCpuTotal(
	studentId uint32,
	lineID int,
	paramType int,
	courseType []int,
	conditions *UserCpuTotalConditions) ([]AtomicUserCpu, error) {

	data := UserCpuTotalBodyStruct{StudentID: studentId, LineID: lineID, Type: paramType, PerPage: 20}
	if paramType == 0 {
		data.Type = 1 //默认为1

	}
	if courseType != nil {
		data.CourseType = courseType
	}

	if conditions != nil {
		data.Conditions = conditions
	}

	body, _ := json.Marshal(data)
	res, err := c.Remote.Post("/v1/ucpu/total", body)
	if err != nil {
		c.Ctx.Log.Warning(fmt.Sprintf("%+v", err))
		return nil, err
	}
	var totalCpu []AtomicUserCpu

	err = json.Unmarshal(res.Data, &totalCpu)
	if err != nil {
		return nil, err
	}
	return totalCpu, nil
}
