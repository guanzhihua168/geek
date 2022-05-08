package repository

import (
	"encoding/json"
	"errors"
	"parent-api-go/global"
	"parent-api-go/pkg/linkerd"
)

type UserRepos struct {
	repository
}

type User struct {
	Id     uint32
	Mobile string
	Email  string
	Active bool
}

type Student struct {
	ID          uint32 `json:"id"`
	ParentID    uint32 `json:"parent_id"`
	No          string `json:"no"`
	Name        string `json:"name"`
	UserType    int    `json:"type"`
	Status      int    `json:"status"`
	Active      int    `json:"active"`
	ClassDevice int    `json:"class_device"`
	StudyStage  int    `json:"study_stage"`
}

func (c *UserRepos) Init() {
	options := linkerdOptions(c.repository)
	options = append(options, linkerd.WithJson())
	c.Remote = linkerd.NewLinkerd(
		global.LinkerdSetting.Host,
		global.LinkerdSetting.AppName,
		global.LinkerdSetting.Token,
		"ucenter",
		options...,
	)
	//c.Remote = linkerd.NewLinkerd(
	//	configs.App.GetString("Linkerd.Host"),
	//	configs.App.GetString("Linkerd.AppName"),
	//	configs.App.GetString("Linkerd.Token"),
	//	"ucenter",
	//	options...,
	//)

}

func (c *UserRepos) ParseAuthToken(token string) uint32 {
	data := struct {
		Token string `json:"token"`
	}{
		token,
	}

	body, _ := json.Marshal(data)
	j, e := c.Remote.Post("/v1/userInfo/getUserInfoByToken", body)
	if e != nil {
		c.Ctx.Log.Warning(e.Error())
		return 0
	}

	u := []User{}
	json.Unmarshal(j.Data, &u)
	if len(u) == 0 {
		return 0
	}
	return u[0].Id
}

// https://yapi.rouchi.com/project/267/interface/api/10414
func (c *UserRepos) getStudentByParents(uids []uint32, column []string) (stu []Student, e error) {
	data := struct {
		ParentIds []uint32 `json:"parent_ids"`
		Columns   []string `json:"columns"`
	}{
		uids,
		column,
	}

	body, _ := json.Marshal(data)
	j, e := c.Remote.Post("/v2/student/getByParentOrNo", body)
	if e != nil {
		c.Ctx.Log.Warning(e.Error())
		return stu, e
	}

	if e = json.Unmarshal(j.Data, &stu); e != nil {
		return stu, e
	}
	return stu, nil
}

func (c UserRepos) GetStuIdByUid(uid uint32) uint32 {
	s, e := c.getStudentByParents([]uint32{uid}, []string{"id"})
	if e != nil || len(s) == 0 {
		return 0
	}
	return s[0].ID
}

func (c *UserRepos) GetLineId(uid uint32) (int, error) {
	atl := AtomicLeadsRepos{}
	c.Ctx.CtxInto(&atl)
	// 根据学生id获取业务线
	l, err := atl.GetLeadsByParentUIDs([]uint32{uid})
	if err != nil {
		return 0, err
	}

	if len(l) == 0 {
		return -1, errors.New("not found  user line id")
	}
	return l[0].LineId, nil
}
