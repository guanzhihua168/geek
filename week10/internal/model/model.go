package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"parent-api-go/global"
	"parent-api-go/pkg/setting"
	"strconv"
	"time"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID     int32 `gorm:"primary_key" json:"id"`
	Active int8  `json:"-"`
}

// 创建数据库连接
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)

	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	//otgorm.AddGormCallbacks(db) //add数据链路
	return db, nil
}

func NewRedisEngine(redisSetting *setting.RedisSettingS) (*redis.Pool, error) {
	RedisPool := &redis.Pool{
		MaxIdle:     redisSetting.MaxIdleConns,
		MaxActive:   redisSetting.MaxOpenConns,
		IdleTimeout: 240 * time.Second,
		// Wait:        true,
		Dial: func() (redis.Conn, error) {
			dbName, _ := strconv.Atoi(redisSetting.Database)
			c, e := redis.Dial(
				"tcp",
				redisSetting.Host+":"+strconv.Itoa(redisSetting.Port),
				redis.DialDatabase(dbName),
				redis.DialPassword(redisSetting.Password),
				redis.DialReadTimeout(time.Duration(redisSetting.TimeOut)*time.Second),
				redis.DialWriteTimeout(time.Duration(redisSetting.TimeOut)*time.Second),
				redis.DialConnectTimeout(time.Duration(redisSetting.TimeOut)*time.Second),
			)
			if e != nil {
				logrus.WithField("redispool", "init").Error(e.Error())
				return nil, e
			}
			return c, e
		},
	}
	return RedisPool, nil
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
