package main

import (
	"flag"
	"gf-example/web-demo/cf"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/skill/conf"
	"github.com/qinchende/gofast/skill/exec"
	"github.com/qinchende/gofast/skill/httpx"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	threadNum     = 10   // 请求的线程数，并发
	totalRequests = 1000 // 总请求数
)

func main() {
	loadConfigDel()
	logx.Info("AutoRequest, I'm running......")

	// 多个线程发起请求
	goWait := sync.WaitGroup{}
	for i := 0; i < threadNum; i++ {
		goWait.Add(1)
		go autoRequest(&goWait)
	}
	goWait.Wait()
	logx.InfoF("Request times: %d, Suc times %d", loopCount, sucCount)
	logx.Info("All threads finished. Now exit. bye bye...")
}

var reduceLog1 *exec.Reduce
var reduceLog2 *exec.Reduce

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func loadConfigDel() {
	var AppCnf cf.ProjectConfig
	var cnfFile = flag.String("f", "../cf/env.yaml", "-f env.[yaml|yml|json]")

	flag.Parse()
	conf.MustLoad(*cnfFile, &AppCnf)

	// reset log config
	pLogConfig := &AppCnf.WebServerCnf.LogConfig
	pLogConfig.AppName += "-autoreq"
	pLogConfig.FileFolder = "../" + pLogConfig.FileFolder

	logx.MustSetup(pLogConfig)
	logx.Info("Hello " + AppCnf.WebServerCnf.AppName + ", config all ready.")
	cf.InitMysql()

	reduceLog1 = exec.NewReduce(time.Second * 5)
	reduceLog2 = exec.NewReduce(time.Second * 30)
}

var (
	loopCount int32 = 0
	sucCount  int32 = 0
)

// Auto Running
func autoRequest(wg *sync.WaitGroup) {
	defer wg.Done()

	for atomic.LoadInt32(&loopCount) < totalRequests {
		lpc := atomic.AddInt32(&loopCount, 1)
		if lpc > totalRequests {
			break
		}

		_, err := httpx.DoRequestGetKV(&httpx.RequestPet{
			ProxyUrl: cf.AppCnf.CurrAppData.ProxyUrl,
			Method:   http.MethodGet,
			Url:      "http://127.0.0.1:8078/request_test_data",
			//QueryArgs: cst.KV{"tok": "t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"},
			//BodyArgs: cst.KV{"tok": "t:NDhDdjdwMEdaWTZoamtnY01o.RALE84mO4YGpAFdPfFEO8gi4NFcvH1kQV9IWmfaJuyc"},
		})

		// do request +++++++++++++++++++++++++++++++++++==
		scc := int32(0)
		if err != nil {
			reduceLog2.DoInterval(lpc == totalRequests, func(skipTimes int32) {
				logx.InfoF("Ret error # %s #, Skip log times: %d", err.Error(), skipTimes)
			})
			// 异常，sleep 1000ms
			time.Sleep(time.Duration(520) * time.Millisecond)
			scc = atomic.LoadInt32(&sucCount)
		} else {
			// 正常，sleep 500ms
			time.Sleep(time.Duration(160) * time.Millisecond)
			scc = atomic.AddInt32(&sucCount, 1)
		}

		reduceLog1.DoInterval(lpc == totalRequests, func(skipTimes int32) {
			logx.InfoF("Request times: %d, Suc times %d, Skip log times: %d", lpc, scc, skipTimes)
		})
	}
}
