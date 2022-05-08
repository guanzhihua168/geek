package global

import (
	"parent-api-go/pkg/logger"
	"parent-api-go/pkg/setting"
)

var (
	ServerSetting                *setting.ServerSettingS
	AppSetting                   *setting.AppSettingS
	EmailSetting                 *setting.EmailSettingS
	JWTSetting                   *setting.JWTSettingS
	DatabaseOldBossMasterSetting *setting.DatabaseSettingS
	DatabaseOldBossSlaveSetting  *setting.DatabaseSettingS
	DatabaseParentMasterSetting  *setting.DatabaseSettingS
	DatabaseParentSlaveSetting   *setting.DatabaseSettingS
	DatabaseMopMasterSetting     *setting.DatabaseSettingS
	DatabaseMopSlaveSetting      *setting.DatabaseSettingS
	Logger                       *logger.Logger
	LinkerdSetting               *setting.LinkerdSettingS
	SkywalkingSetting            *setting.SkywalkingSettingS
	RabbitmqSetting              *setting.RabbitmqSettingS
	RedisDefaultSetting          *setting.RedisSettingS
	AppContext                   = "AppContext" //从uggo-service移植过来
)
