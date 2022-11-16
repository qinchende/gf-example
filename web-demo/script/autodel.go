package main

import (
	"flag"
	"gf-example/web-demo/cf"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/skill/conf"
	"github.com/qinchende/gofast/skill/exec"
	"github.com/qinchende/gofast/skill/httpx"
	"net/http"
	"time"
)

func main() {
	loadConfigDel()
	logx.Info("AutoDel, I'm running......")
	autoDelRecords()
}

var reduceLog1 *exec.Reduce
var reduceLog2 *exec.Reduce

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func loadConfigDel() {
	var AppCnf cf.ProjectConfig
	var cnfFile = flag.String("f", "../cf/env.yaml", "-f env.[yaml|yml|json]")

	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)
	logx.MustSetup(&AppCnf.WebServerCnf.LogConfig)
	logx.Info("Hello " + AppCnf.WebServerCnf.AppName + ", config all ready.")
	cf.InitMysql()

	reduceLog1 = exec.NewReduce(time.Second * 5)
	reduceLog2 = exec.NewReduce(time.Second * 30)
}

// Auto Running
func autoDelRecords() {
	count := 0
	for count < 100000 {
		count++
		reduceLog1.DoOrNot(func(skipTimes int32) {
			logx.InfoF("Run times: %d, Skip log times: %d", count, skipTimes)
		})
		doOneRequest()
		//time.Sleep(30 * time.Millisecond)
	}
}

func doOneRequest() {
	_, err := httpx.DoRequestGetKV(&httpx.RequestPet{
		ProxyUrl: cf.AppCnf.CurrAppData.ProxyUrl,
		Method:   http.MethodGet,
		Url:      "http://127.0.0.1:8078/request_test_data",
		//QueryArgs: cst.KV{"tok": "t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"},
		//BodyArgs: cst.KV{"tok": "t:NDhDdjdwMEdaWTZoamtnY01o.RALE84mO4YGpAFdPfFEO8gi4NFcvH1kQV9IWmfaJuyc"},
	})
	if err != nil {
		reduceLog2.DoOrNot(func(skipTimes int32) {
			logx.InfoF("Ret error %s, Skip log times: %d", err.Error(), skipTimes)
		})
	}
}
