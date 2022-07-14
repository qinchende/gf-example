package cf

import (
	"flag"
	"github.com/qinchende/gofast/connx/gform"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/fstx"
	"github.com/qinchende/gofast/jwtx"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/skill/conf"
)

type AppConfigEntity struct {
	WebServerCnf   fst.GfConfig    `v:"required"`
	SdxSessCnf     jwtx.SdxSessCnf `v:"required"`
	MysqlGoZeroCnf gform.ConnCnf   `v:"required"`
}

var AppCnf AppConfigEntity
var cnfFile = flag.String("f", "cf/env.yaml", "-f env.[yaml|yml|json]")

func InitEnvConfig() {
	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)
	fstx.InitLogger(&AppCnf.WebServerCnf.LogConfig)
	logx.Info("Hello " + AppCnf.WebServerCnf.Name + ", Welcome.")

	// initGoRedis()
	// tryGoRedis()
	initRedisSession()
	initMysql()
	initGorm()
}
