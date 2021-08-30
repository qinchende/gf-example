package config

import (
	"fmt"
	"github.com/qinchende/gofast/connx/redis"
	"github.com/qinchende/gofast/jwtx"
)

var RedisA *redis.GoRedisX

func initGoRedis() {
	RedisA = redis.NewGoRedis(&SysCnf.SdxSessCnf.RedisConnCnf)
}

func tryGoRedis() {
	pong, err := RedisA.Ping()

	if err != nil {
		fmt.Println("Ping failed", err)
	} else {
		fmt.Printf("Ping val is %s", pong)
	}
}

// init sdx session with redis store
func initRedisSession() {
	sdxSess := jwtx.SdxSession{
		SdxSessConfig: SysCnf.SdxSessCnf,
	}
	sdxSess.Init()
}
