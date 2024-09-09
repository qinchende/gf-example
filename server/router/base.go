package router

import (
	"gf-example/server/router/fit"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/sdx"
)

func LoadRoutes(app *fst.GoFast) {
	// 1. 演示添加第一级中间件，第一级中间件不带fst.Context，而是原始的[req, res]
	app.UseHttpHandler(fit.MyFitDemo)

	// 2. 全局中间件链标准组合（拦截器）
	app.UseGlobal(sdx.SuperHandlers) // 框架自带闪电侠超级中间件链

	// 3. 根路由，中间件。
	// Note: 匹配到路由之后开始走这里的逻辑，执行过滤器
	app.Before(sdx.PmsParser) // 解析请求参数，构造 ctx.Pms
	app.AfterMatch()          // 匹配路由之后，开始执行中间件链之前

	// 4. url api routes lists
	apiRoutes(app)
}
