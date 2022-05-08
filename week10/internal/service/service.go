package service

import (
	"context"
	"parent-api-go/global"
	"parent-api-go/internal/dao"
	"parent-api-go/internal/repository"
)

type Service struct {
	ctx              context.Context
	daoOldBossMaster *dao.Dao
	daoOldBossSlave  *dao.Dao
	daoParentMaster  *dao.Dao
	daoParentSlave   *dao.Dao
	daoMopMaster     *dao.Dao
	daoMopSlave      *dao.Dao
}

//type RequestDeviceP struct {
//	Akey    string `json:"_aKey"`
//	Channel string `json:"_channel"`
//	Did     string `json:"_dId"`
//	Lang    string `json:"_lang"`
//	Module  string `json:"_module"`
//	Nid     string `json:"_nId"`
//	Pname   string `json:"_pName"`
//	T       string `json:"_t"`
//	Udid    string `json:"_udid"`
//	VApp    string `json:"_vApp"`
//	VName   string `json:"_vName"`
//	Vos     string `json:"_vOs"`
//}
//type RequestDeviceP struct {
//	VName string `json:"_vName"`
//	PName string `json:"_pName"`
//	AKey string `json:"_aKey"`
//	T int `json:"_t"`
//	NID string `json:"_nId"`
//	VOs string `json:"_vOs"`
//	Lang string `json:"_lang"`
//	Udid string `json:"_udid"`
//	VApp string `json:"_vApp"`
//	DID string `json:"_dId"`
//}

type RequestDeviceP struct {
	AKey  string `json:"_aKey"`
	VName string `json:"_vName"`
	PName string `json:"_pName"`
	Lang  string `json:"_lang"`
	VApp  string `json:"_vApp"`
	VOs   string `json:"_vOs"`
	NID   string `json:"_nId"`
	DID   string `json:"_dId"`
	UDID  string `json:"_udid"`
	//T string `json:"_t"`//安卓与ios转的类型不一致（iOS是int， 安卓是字符串）先不接收
	Channel string `json:"_channel"`
	Module  string `json:"_module"`
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}

	//svc.daoOldBossMaster = dao.New(otgorm.WithContext(svc.ctx,global.OldBossMasterDBEngine))
	//svc.daoOldBossSlave = dao.New(otgorm.WithContext(svc.ctx, global.OldBossSlaveDBEngine))
	//svc.daoParentMaster = dao.New(otgorm.WithContext(svc.ctx, global.ParentMasterDBEngine))
	//svc.daoParentSlave = dao.New(otgorm.WithContext(svc.ctx, global.ParentSlaveDBEngine))
	//svc.daoMopMaster = dao.New(otgorm.WithContext(svc.ctx, global.MopMasterDBEngine))
	//svc.daoMopSlave = dao.New(otgorm.WithContext(svc.ctx, global.MopSlaveDBEngine))
	svc.daoOldBossMaster = dao.New(global.OldBossMasterDBEngine)
	svc.daoOldBossSlave = dao.New(global.OldBossSlaveDBEngine)
	svc.daoParentMaster = dao.New(global.ParentMasterDBEngine)
	svc.daoParentSlave = dao.New(global.ParentSlaveDBEngine)
	svc.daoMopMaster = dao.New(global.MopMasterDBEngine)
	svc.daoMopSlave = dao.New(global.MopSlaveDBEngine)

	return svc
}

/*
	获取当前家长的所属于的业务线，如果没有取到则默认为美标
	todo 加cache
*/
func (svc *Service) GetParentLineId(uid uint32) (uint8, error) {
	parentLineId := 0
	leadsRepos := repository.AtomicLeadsRepos{}
	leadsRepos.Init()
	userLeads, err := leadsRepos.GetLeadsByParentUIDs([]uint32{uid})
	if len(userLeads) == 0 || err != nil {
		return 0, err
	}

	parentLineId = userLeads[0].LineId
	if parentLineId == 0 {
		parentLineId = global.AmeLineId
	}

	return uint8(parentLineId), nil
}

func (svc *Service) GetStudentId(uid uint32) uint32 {
	userRepos := repository.UserRepos{}
	userRepos.Init()
	return userRepos.GetStuIdByUid(uid)
}
