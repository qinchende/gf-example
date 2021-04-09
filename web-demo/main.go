package main

import (
	"gf-example/web-demo/config"
	"gf-example/web-demo/route"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

// 应用启动前设置全局参数
func init() {
	// log.SetPrefix("[GoFast]")    // 前置字符串加上特定标记
	// log.SetFlags(log.Lmsgprefix) // 取消前置字符串
	// log.SetFlags(log.LstdFlags)  // 设置成日期+时间 格式
}

func main() {
	// TODO: 1. 初始化配置，连接数据库，创建Server
	config.InitEnvConfig()
	serverCnf := &config.EnvParams.WebServerCnf
	serverCnf.Name = "Gf-Web-Demo"
	app := fst.CreateServer(serverCnf)

	// TODO：2. 加载中间件、路由
	route.LoadRoutes(app)

	// TODO: 3. 启动Server, 等待请求
	// 开始监听接收请求
	logx.Infof("Listening and serving HTTP on %s\n", app.Addr)
	lisErr := app.Listen(app.Addr)
	if lisErr != nil {
		logx.Error(lisErr)
	}
}
