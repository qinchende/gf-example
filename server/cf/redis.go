package cf

import (
	"github.com/qinchende/gofast/connx/redis"
	"github.com/qinchende/gofast/sdx"
)

var DRedisTask *redis.GfRedis // 记录任务信息

// init sdx session with redis store
func initRedisForSession() {
	DRedisTask = redis.NewGoRedis(&Cnf.SessionCnf.RedisConn)
	sdx.SetSessionDB(&sdx.SessionDB{
		SessionConfig: Cnf.SessionCnf,
	})
}
