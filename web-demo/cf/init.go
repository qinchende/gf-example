package cf

import (
	"flag"
	"github.com/qinchende/gofast/connx/orm"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
	"github.com/qinchende/gofast/skill/conf"
)

type ProjectConfig struct {
	WebServerCnf   fst.GfConfig     `v:"required"`
	SessionCnf     sdx.SessionCnf   `v:"required"`
	MysqlGoZeroCnf orm.MysqlConnCnf `v:"required"`
	CurrAppParams  appParams        `v:"required"`
}

var AppCnf ProjectConfig

func MustAppConfig() {
	var cnfFile = flag.String("f", "cf/env.yaml", "-f env.[yaml|yml|json]")
	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)

	logConfig := &AppCnf.WebServerCnf.LogConfig
	logConfig.AppName = AppCnf.WebServerCnf.AppName
	logConfig.ServerName = AppCnf.WebServerCnf.ServerName
	logx.MustSetup(logConfig)
	logx.Info("Hello " + logConfig.AppName + ", config data loaded.")

	initAppParams()
	initRedisForSession()
	InitMysql()
}
