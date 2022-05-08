package repository

import (
	"encoding/json"
	"parent-api-go/global"
	"parent-api-go/pkg/linkerd"
)

type UserLabelRepos struct {
	repository
}

type UserLabels struct {
	TagId int32
}

func (c *UserLabelRepos) Init() {
	options := linkerdOptions(c.repository)
	options = append(options, linkerd.WithJson())
	c.Remote = linkerd.NewLinkerd(
		global.LinkerdSetting.Host,
		global.LinkerdSetting.AppName,
		global.LinkerdSetting.Token,
		"service_userlabel", // 用户标签服务
		options...,
	)
}

// 获取指定用户标签
func (c *UserLabelRepos) GetUserLabel(userId uint32) (ul []*UserLabels, e error) {
	data := struct {
		UserId uint32 `json:"user_id"`
	}{userId}

	body, _ := json.Marshal(data)
	j, e := c.Remote.Post("/V1/Label/UserLabel/getUserLabel", body)
	if e != nil {
		c.Ctx.Log.Warning(e.Error())
		return
	}

	e = json.Unmarshal(j.Data, &ul)
	if e != nil {
		c.Ctx.Log.Warning("[GetUserLabel]json.Unmarshal error:" + e.Error())
		return
	}

	return
}
