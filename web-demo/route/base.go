package route

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
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

	// 根路由 特殊情况处理, 不写的话就是默认处理函数
	app.NoRoute(func(c *fst.Context) {
		c.String(http.StatusNotFound, "Custom NoRoute func -> 404-Can't find the path.")
	})
	app.NoMethod(func(c *fst.Context) {
		c.String(http.StatusMethodNotAllowed, "Custom NoMethod func -> 405-Method not allowed.")
	})

	// 2.1. 全局中间件（拦截器）
	// Note: 请求进来，并没有定位到具体的路由。就需要走这些过滤器
	// 所有的请求都要走这里指定的拦截器，发生错误就直接中断返回
	app.UseGlobal(sdx.DefGlobalFits) // 默认的一组中间件
	//app.UseGlobalFit(filter.MyFitDemo) // 自定义顶层中间件
	// 2.2. 全局，第二级中间件
	app.UseGlobal(sdx.DefGlobalHandlers)

	// 3. 根路由，中间件。
	// Note: 匹配到路由之后开始走这里的逻辑，执行过滤器
	app.Before(sdx.PmsParser)   // 解析请求参数，构造 ctx.Pms
	app.Before(sdx.SessBuilder) // “闪电侠Session”：所有请求要携带tok信息，没有就自动分配一个

	// 4. all routes lists
	apiRoutes(app)
}
