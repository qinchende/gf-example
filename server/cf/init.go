package cf

import (
	"flag"
	"github.com/qinchende/gofast/aid/conf"
	"github.com/qinchende/gofast/aid/logx"
	"github.com/qinchende/gofast/connx/orm"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/sdx"
)

type AppConfig struct {
	ServerCnf     fst.ServerConfig  `v:"must"`
	SdxMidCnf     sdx.MidConfig     `v:"must"`
	SessionCnf    sdx.SessionConfig `v:"must"`
	MysqlDemoCnf  orm.MysqlConfig   `v:"must"`
	CurrAppParams appParams         `v:"must"`
}

var Cnf AppConfig

func MustInitConfig() {
	var cnfFile = flag.String("f", "cf/conf.yaml", "-f conf.[yaml|yml|json]")
	flag.Parse()
	conf.MustLoad(*cnfFile, &Cnf)

	// 最先初始化日志系统
	logCnf := &Cnf.ServerCnf.LogConfig
	logCnf.AppName = Cnf.ServerCnf.AppName
	logCnf.HostName = Cnf.ServerCnf.HostName
	logx.MustSetup(logCnf)
	logx.Info("Hello " + logCnf.AppName + ", config data loaded.")

	// 初始化中间件控制参数
	sdx.SetMidConfig(&Cnf.SdxMidCnf)

	initAppParams()
	initRedisForSession()
	InitMysql()
}
