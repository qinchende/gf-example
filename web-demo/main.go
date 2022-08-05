package main

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/route"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func main() {
	// Step1. 初始化配置，连接数据库，创建Server
	cf.MustAppConfig()
	app := fst.CreateServer(&cf.AppCnf.WebServerCnf)

	// Step2. 加载中间件、路由
	route.LoadRoutes(app)

	logx.Stack("this is stack log")
	logx.Slow("this is slow log")
	logx.Stat("this is stat log")

	// Step3. 启动Server Listen, 等待请求
	logx.InfoF("Listening and serving HTTP on %s", app.ListenAddr)
	if lisErr := app.Listen(); lisErr != nil {
		logx.ErrorFatal(lisErr)
	}
}
