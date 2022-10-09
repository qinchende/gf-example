package cf

import (
	"flag"
	"github.com/qinchende/gofast/connx/gform"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
	"github.com/qinchende/gofast/skill/conf"
)

type AppConfigEntity struct {
	WebServerCnf   fst.GfConfig     `v:"required"`
	RedisSessCnf   sdx.RedisSessCnf `v:"required"`
	MysqlGoZeroCnf gform.ConnCnf    `v:"required"`
}

var AppCnf AppConfigEntity

func MustAppConfig() {
	var cnfFile = flag.String("f", "cf/env.yaml", "-f env.[yaml|yml|json]")
	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)
	logx.MustSetup(&AppCnf.WebServerCnf.LogConfig)
	logx.Info("Hello " + AppCnf.WebServerCnf.AppName + ", config all ready.")

	initRedisForSession()
	InitMysql()
}
