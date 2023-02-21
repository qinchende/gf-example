package main

import (
	"flag"
	"gf-example/web-demo/cf"
	"github.com/qinchende/gofast/cst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/skill/conf"
	"github.com/qinchende/gofast/skill/exec"
	"github.com/qinchende/gofast/skill/httpx"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// 测试熔断降载的实际表现。
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
	time.Sleep(2 * time.Second) // 确保退出之前打印日志
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
	pLogConfig.LogMedium = "console"

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

var randTool = rand.New(rand.NewSource(time.Now().UnixNano()))

// Auto Running
func autoRequest(wg *sync.WaitGroup) {
	defer wg.Done()

	for atomic.LoadInt32(&loopCount) < totalRequests {
		lpc := atomic.AddInt32(&loopCount, 1)
		if lpc > totalRequests {
			break
		}

		// 主动设置请求处理的耗时时间
		var delayMS int32 = 0
		if lpc < 33 {
			delayMS += lpc * 3
		} else if lpc < 200 {
			delayMS += (80 + randTool.Int31n(20))
		} else if lpc < 400 {
			delayMS += (80 + randTool.Int31n(200))
		} else if lpc < 600 {
			delayMS += (50 + randTool.Int31n(100))
		} else if lpc < 700 {
			delayMS += (700 - lpc)
		} else {
			delayMS = 10
		}

		_, err := httpx.DoRequestGetKV(&httpx.RequestPet{
			ProxyUrl:  cf.DParams.ProxyUrl,
			Method:    http.MethodGet,
			Url:       "http://127.0.0.1:8078/request_test_data",
			QueryArgs: cst.KV{"Count": lpc, "DelayMS": delayMS},
			//QueryArgs: cst.KV{"tok": "t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"},
			//BodyArgs: cst.KV{"tok": "t:NDhDdjdwMEdaWTZoamtnY01o.RALE84mO4YGpAFdPfFEO8gi4NFcvH1kQV9IWmfaJuyc"},
		})

		// do request +++++++++++++++++++++++++++++++++++==
		scc := int32(0)
		if err != nil {
			reduceLog2.DoInterval(lpc == totalRequests, func(skipTimes int32) {
				logx.InfoF("Ret error # %s #, Skip log times: %d", err.Error(), skipTimes)
			})
			// 异常
			time.Sleep(time.Duration(500) * time.Millisecond)
			scc = atomic.LoadInt32(&sucCount)
		} else {
			// 正常
			time.Sleep(time.Duration(500) * time.Millisecond)
			scc = atomic.AddInt32(&sucCount, 1)
		}

		reduceLog1.DoInterval(lpc == totalRequests, func(skipTimes int32) {
			logx.InfoF("Request times: %d, Suc times %d, Skip log times: %d", lpc, scc, skipTimes)
		})
	}
}
