package config

import (
	"fmt"
	"github.com/qinchende/gofast/connx/redis"
	"github.com/qinchende/gofast/jwtx"
)

var RedisA *redis.GoRedisX

func initGoRedis() {
	RedisA = redis.NewGoRedis(&EnvParams.SdxSessCnf.RedisConnCnf)
}

func tryGoRedis() {
	pong, err := RedisA.Ping()

	if err != nil {
		fmt.Println("Ping failed", err)
	} else {
		fmt.Printf("Ping val is %s\n", pong)
	}
}

// 初始化 sdx session redis
func initRedisSession() {
	sdxSess := jwtx.SdxSession{
		SdxSessConfig: EnvParams.SdxSessCnf,
	}
	sdxSess.Init()
}
