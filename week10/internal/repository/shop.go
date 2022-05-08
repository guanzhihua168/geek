package repository

import (
	"encoding/json"
	"parent-api-go/global"
	"parent-api-go/pkg/linkerd"
)

type ShopRepos struct {
	repository
}
type Balance struct {
	Coin int
}

func (s *ShopRepos) Init() {
	options := linkerdOptions(s.repository)
	options = append(options, linkerd.WithJson())
	s.Remote = linkerd.NewLinkerd(
		global.LinkerdSetting.Host,
		global.LinkerdSetting.AppName,
		global.LinkerdSetting.Token,
		"service_course",
		options...,
	)
}

func (s *ShopRepos) GetUserCoinBalance(userId uint32) int {
	s.Init()

	data := struct {
		UserID uint32 `json:"user_id"`
	}{
		userId,
	}

	type BalanceResult []int

	body, _ := json.Marshal(data)
	j, e := s.Remote.Post("/V2/Coin/getUserCoin", body)
	if e != nil {
		s.Ctx.Log.Warning(e.Error())
		return 0
	}

	r := BalanceResult{}
	_ = json.Unmarshal(j.Data, &r)
	if len(r) == 0 {
		return 0
	}
	return r[0]
}
