package route

import (
	"gf-example/web-demo/route/mid"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/fstx"
	"github.com/qinchende/gofast/jwtx"
	"github.com/qinchende/gofast/logx"
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
	//app.NoRoute(func(ctx *fst.Context) {
	//	ctx.String(http.StatusNotFound, "404-Can't find the path.")
	//})
	//app.NoMethod(func(ctx *fst.Context) {
	//	ctx.String(http.StatusMethodNotAllowed, "405-Method not allowed.")
	//})

	// 2.1. 全局中间件（拦截器）
	// Note: 请求进来，并没有定位到具体的路由。就需要走这些过滤器
	// 所有的请求都要走这里指定的拦截器，发生错误就直接中断返回
	app.Apply(fstx.DefaultFits) // 默认的一组中间件
	app.Fit(mid.MyFitDemo)      // 自定义中间件
	// 2.2. 全局，第二级中间件
	app.Apply(fstx.DefaultHandlers)

	// 3. 根路由，中间件。
	// Note: 匹配到路由之后开始走这里的逻辑，执行过滤器
	app.Before(fstx.PmsParser)      // 解析请求参数，构造 ctx.Pms
	app.Before(jwtx.SdxSessBuilder) // “闪电侠Session”：所有请求要携带tok信息，没有就自动分配一个

	// 4. all routes lists
	routesList(app)
}
