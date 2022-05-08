package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string        // 自定义键值对名称
	FillInterval time.Duration //发方令牌的间隔时间
	Capacity     int64         //令牌桶容量
	Quantum      int64         //每闪达到间隔时间后，发放具体令牌桶的数量
}
