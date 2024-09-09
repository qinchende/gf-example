package cf

import (
	"flag"
	"github.com/qinchende/gofast/aid/conf"
	"github.com/qinchende/gofast/aid/logx"
	"github.com/qinchende/gofast/connx/orm"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/sdx"
)

type ProjectConfig struct {
	WebServerCnf  fst.AppConfig    `v:"must"`
	SessionCnf    sdx.SessionCnf   `v:"must"`
	MysqlDemoCnf  orm.MysqlConnCnf `v:"must"`
	CurrAppParams appParams        `v:"must"`
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
