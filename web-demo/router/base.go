package router

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
)

func LoadRoutes(app *fst.GoFast) {
	// 1. 基础
	// 应用级事件
	app.OnReady(func(fast *fst.GoFast) {
		logx.Info("App OnReady Call.")
	})
	app.OnClose(func(fast *fst.GoFast) {
		logx.Info("App OnClose Call.")
	})

	// 2. 全局中间件链（拦截器）
	app.UseGlobal(sdx.SuperHandlers) // 框架自带闪电侠超级中间件链

	// 3. 根路由，中间件。
	// Note: 匹配到路由之后开始走这里的逻辑，执行过滤器
	app.Before(sdx.PmsParser) // 解析请求参数，构造 ctx.Pms
	//app.Before(sdx.SessBuilder) // “闪电侠Session”：所有请求要携带tok信息，没有就自动分配一个

	// 4. all routes lists
	apiRoutes(app)
}
