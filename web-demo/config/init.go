package config

import (
	"flag"
	"github.com/qinchende/gofast/connx/mysql"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/fstx"
	"github.com/qinchende/gofast/jwtx"
	"github.com/qinchende/gofast/skill/conf"
)

type SysConfigEntity struct {
	WebServerCnf fst.AppConfig      `cnf:",NA"`
	SdxSessCnf   jwtx.SdxSessConfig `cnf:",NA"`
	SqlGoZeroCnf mysql.ConnConfig   `cnf:",NA"`
}

var SysCnf SysConfigEntity
var cnfFile = flag.String("f", "config/env.yaml", "-f env.[yaml|yml|json]")

func InitEnvConfig() {
	flag.Parse()
	conf.MustLoad(*cnfFile, &SysCnf)
	fstx.InitLogger(&SysCnf.WebServerCnf.LogConfig)

	// initGoRedis()
	// tryGoRedis()
	initRedisSession()
	//initMysql()
	initGorm()
}
