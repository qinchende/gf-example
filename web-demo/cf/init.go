package cf

import (
	"flag"
	"github.com/qinchende/gofast/connx/gform"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/fstx"
	"github.com/qinchende/gofast/jwtx"
	"github.com/qinchende/gofast/skill/conf"
)

type AppConfigEntity struct {
	WebServerCnf   fst.AppConfig   `cnf:",NA"`
	SdxSessCnf     jwtx.SdxSessCnf `cnf:",NA"`
	MysqlGoZeroCnf gform.ConnCnf   `cnf:",NA"`
}

var AppCnf AppConfigEntity
var cnfFile = flag.String("f", "cf/env.yaml", "-f env.[yaml|yml|json]")

func InitEnvConfig() {
	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)
	fstx.InitLogger(&AppCnf.WebServerCnf.LogConfig)

	// initGoRedis()
	// tryGoRedis()
	initRedisSession()
	initMysql()
	initGorm()
}
