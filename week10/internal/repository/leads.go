package repository

import (
	"encoding/json"
	"fmt"
	"parent-api-go/global"
	"parent-api-go/pkg/linkerd"
)

type AtomicLeadsRepos struct {
	repository
}

func (c *AtomicLeadsRepos) Init() {
	options := linkerdOptions(c.repository)
	options = append(options, linkerd.WithJson())
	c.Remote = linkerd.NewLinkerd(
		global.LinkerdSetting.Host,
		global.LinkerdSetting.AppName,
		global.LinkerdSetting.Token,
		"atomic_leads",
		options...,
	)
}

type UserLead struct {
	UserId uint `json:"user_id"`
	LineId int  `json:"line_id"`
	Active int  `json:"active"`
}

//通过家长用户ID获取用户属于那个业务线
//https://yapi.rouchi.com/project/384/interface/api/10734
func (c *AtomicLeadsRepos) GetLeadsByParentUIDs(userIDs []uint32) (ul []UserLead, err error) {
	type BodyStruct struct {
		UserIDs     []uint32 `json:"user_ids"`
		QueryFields []string `json:"query_fields"`
		LineID      int      `json:"line_id"`
	}
	data := BodyStruct{
		UserIDs:     userIDs,
		QueryFields: []string{"user_id", "line_id", "active"},
		LineID:      0,
	}

	body, _ := json.Marshal(data)
	j, e := c.Remote.Post("/v1/leads/leadsinfo/getLeadsByUserIds", body)
	if e != nil {
		c.Ctx.Log.Warning(fmt.Sprintf("%+v", e))
		return nil, err
	}

	err = json.Unmarshal(j.Data, &ul)
	if err != nil {
		return nil, err
	}
	return ul, err

}
