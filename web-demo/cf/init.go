package cf

import (
	"flag"
	"github.com/qinchende/gofast/connx/gform"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
	"github.com/qinchende/gofast/skill/conf"
)

type ProjectConfig struct {
	WebServerCnf   fst.GfConfig     `v:"required"`
	RedisSessCnf   sdx.RedisSessCnf `v:"required"`
	MysqlGoZeroCnf gform.ConnCnf    `v:"required"`
	CurrAppParams  appParams        `v:"required"`
}

var AppCnf ProjectConfig

func MustAppConfig() {
	var cnfFile = flag.String("f", "cf/env.yaml", "-f env.[yaml|yml|json]")
	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)

	logConfig := &AppCnf.WebServerCnf.LogConfig
	logConfig.AppName = AppCnf.WebServerCnf.AppName
	logConfig.ServerNo = AppCnf.WebServerCnf.ServerNo
	logx.MustSetup(logConfig)
	logx.Info("Hello " + logConfig.AppName + ", config data loaded.")

	initAppParams()
	initRedisForSession()
	InitMysql()
}
