package cf

import (
	"github.com/qinchende/gofast/sdx"
)

//var RedisA *gfrds.GfRedis

//
//func initGoRedis() {
//	RedisA = gfrds.NewGoRedis(&AppCnf.SdxSessCnf.RedisConnCnf)
//}
//
//func tryGoRedis() {
//	pong, err := RedisA.Ping()
//
//	if err != nil {
//		fmt.Println("Ping failed", err)
//	} else {
//		fmt.Printf("Ping val is %s", pong)
//	}
//}

// init sdx session with redis store
func initRedisForSession() {
	sdx.SetupSession(&sdx.RedisSessionDB{
		RedisSessCnf: AppCnf.RedisSessCnf,
	})
}
