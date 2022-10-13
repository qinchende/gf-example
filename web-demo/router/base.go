package router

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
	"github.com/qinchende/gofast/sdx/mid"
	"net/http"
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

	// 2.1. 匹配路由中间件（拦截器）
	//app.UseHttpHandler(fit.MyFitDemo) // 自定义顶层中间件
	app.UseGlobal(sdx.DefHandlers) // 框架默认中间件链
	// 2.2. 特殊路由中间件
	specialRoutes(app)

	// 3. 根路由，中间件。
	// Note: 匹配到路由之后开始走这里的逻辑，执行过滤器
	app.Before(sdx.PmsParser)   // 解析请求参数，构造 ctx.Pms
	app.Before(sdx.SessBuilder) // “闪电侠Session”：所有请求要携带tok信息，没有就自动分配一个

	// 4. all routes lists
	apiRoutes(app)
}

// 正确匹配路由之外的情况，比如特殊的404,504等路由处理链
func specialRoutes(app *fst.GoFast) {
	app.SpecialBefore(mid.LoggerMini)

	// 根路由 特殊情况处理, 不写的话就是默认处理函数
	app.Reg404(func(c *fst.Context) {
		c.AbortDirect(http.StatusNotFound, "Custom NoRoute func -> 404-Can't find the path.")
	})
	app.Reg405(func(c *fst.Context) {
		c.AbortDirect(http.StatusMethodNotAllowed, "Custom NoMethod func -> 405-Method not allowed.")
	})
}
