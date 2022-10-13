package main

import (
	"flag"
	"gf-example/web-demo/cf"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/skill/conf"
	"time"
)

func main() {
	loadConfigDel()
	logx.Info("AutoDel, I'm running......")
	autoDelRecords()
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func loadConfigDel() {
	var AppCnf cf.ProjectConfig
	var cnfFile = flag.String("f", "../cf/env.yaml", "-f env.[yaml|yml|json]")

	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)
	logx.MustSetup(&AppCnf.WebServerCnf.LogConfig)
	logx.Info("Hello " + AppCnf.WebServerCnf.AppName + ", config all ready.")
	cf.InitMysql()
}

// Auto Running
func autoDelRecords() {
	count := 0
	for true {
		count++
		logx.InfoF("Run times: %d\n", count)
		time.Sleep(60 * time.Second)
	}
}
