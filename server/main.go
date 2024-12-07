package main

import (
	"gf-example/server/cf"
	"gf-example/server/router"
	"github.com/qinchende/gofast/aid/logx"
	"github.com/qinchende/gofast/fst"
)

func main() {
	cf.MustInitConfig()                        // Step1. 初始化配置，连接数据库等
	app := fst.CreateServer(&cf.Cnf.ServerCnf) // Step2. 创建 APP Server
	app.Apply(router.LoadRoutes)               // Step3. 加载中间件、路由
	app.Listen()                               // Step4. 启动Server Listen, 等待请求
	logx.Info().SendMsg("Listening on " + app.ListenAddr)
}

//func main() {
//	MainSample()
//}
