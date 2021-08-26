package config

import (
	"flag"
	"github.com/qinchende/gofast/connx/mysql"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/fstx"
	"github.com/qinchende/gofast/jwtx"
	"github.com/qinchende/gofast/skill/conf"
)

type YmlConfig struct {
	WebServerCnf fst.AppConfig      `json:",optional"`
	SdxSessCnf   jwtx.SdxSessConfig `json:",optional"`
	SqlGoZeroCnf mysql.ConnConfig   `json:",optional"`
}

var EnvParams YmlConfig
var cfgFile = flag.String("f", "config/env.yaml", "-f env.[yaml|yml|json]")

func InitEnvConfig() {
	flag.Parse()
	conf.MustLoad(*cfgFile, &EnvParams)
	fstx.InitLogger(&EnvParams.WebServerCnf.LogConfig)

	// initGoRedis()
	// tryGoRedis()
	initRedisSession()
	//initMysql()
	initGorm()
}
