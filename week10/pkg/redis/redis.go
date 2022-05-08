package redis

import (
	"parent-api-go/global"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

var RedisPool *redis.Pool
var RedisPrefix string

type RedisGo struct {
	Prefix string
}

func (c *RedisGo) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	co := global.RedisPoolDefaultEngine.Get()
	defer co.Close()
	reply, err = c.do(co, commandName, args...)
	return
}

func (c *RedisGo) Doing(co redis.Conn, commandName string, args ...interface{}) (reply interface{}, err error) {
	reply, err = c.do(co, commandName, args...)
	return
}

func (c RedisGo) wrapKeyName(key string) string {
	if len(c.Prefix) > 0 {
		key = c.Prefix + ":" + key
	}
	key = "service_activity:" + key
	return key
}

func (c RedisGo) do(co redis.Conn, commandName string, args ...interface{}) (reply interface{}, err error) {

	params := []interface{}{}
	params = append(params, args...)
	params[0] = c.wrapKeyName(params[0].(string))

	reply, e := co.Do(commandName, params...)
	if e != nil {
		logrus.WithFields(logrus.Fields{
			"type": "redis go",
			"cmd":  commandName,
			"args": params,
		}).Warning(e.Error())
	}
	return
}

func (c RedisGo) GetConn() redis.Conn {
	return global.RedisPoolDefaultEngine.Get()
}

// 分布式锁 支持过期时间
func (c RedisGo) Mutex(name string, expiry time.Duration) *redsync.Mutex {
	rsn := redsync.New(redigo.NewPool(global.RedisPoolDefaultEngine)).NewMutex(
		c.wrapKeyName(name),
		redsync.WithExpiry(expiry),
		redsync.WithTries(1),
	)
	return rsn
}
