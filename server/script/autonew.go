package main

import (
	"flag"
	"gf-example/server/cf"
	"github.com/qinchende/gofast/aid/conf"
	"github.com/qinchende/gofast/aid/logx"
	"time"
)

func main() {
	loadConfigNew()
	logx.Info("AutoNew, I'm running......")
	autoCreateRecords()
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func loadConfigNew() {
	var AppCnf cf.ProjectConfig
	var cnfFile = flag.String("f", "../cf/env.yaml", "-f env.[yaml|yml|json]")

	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)
	logx.MustSetup(&AppCnf.WebServerCnf.LogConfig)
	logx.Info("Hello " + AppCnf.WebServerCnf.AppName + ", config all ready.")
	cf.InitMysql()
}

// Auto Running
func autoCreateRecords() {
	count := 0
	for true {
		count++
		logx.InfoF("Run times: %d\n", count)
		time.Sleep(60 * time.Second)
	}
}
