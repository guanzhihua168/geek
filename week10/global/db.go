package global

import "github.com/jinzhu/gorm"

var (
	OldBossMasterDBEngine *gorm.DB
	OldBossSlaveDBEngine  *gorm.DB
	ParentMasterDBEngine  *gorm.DB
	ParentSlaveDBEngine   *gorm.DB
	MopMasterDBEngine     *gorm.DB
	MopSlaveDBEngine      *gorm.DB
)
