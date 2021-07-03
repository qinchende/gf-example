package route

import (
	"gf-example/web-demo/route/auth"
	"gf-example/web-demo/route/ghost"
	"gf-example/web-demo/route/mid"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/fstx"
	"github.com/qinchende/gofast/jwtx"
	"github.com/qinchende/gofast/logx"
	"net/http"
)

func serverSetup(app *fst.GoFast) {
	// 应用级事件
	app.OnReady(func(fast *fst.GoFast) {
		logx.Info("App OnReady Call.")
	})
	app.OnClose(func(fast *fst.GoFast) {
		logx.Info("App OnClose Call.")
	})

	// 根路由 特殊情况处理 +++++++++++++++++++++++++++++
	app.NoRoute(func(ctx *fst.Context) {
		ctx.JSON(http.StatusNotFound, "404-Can't find the path.")
	})
	app.NoMethod(func(ctx *fst.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, "405-Method not allowed.")
	})
}

func LoadRoutes(app *fst.GoFast) {
	// 1. 基础
	serverSetup(app)

	// 2. 全局中间件（拦截器）
	// Note: 请求进来，并没有定位到具体的路由。就需要走这些过滤器
	// 所有的请求都要走这里指定的拦截器，发生错误就直接中断返回
	app.InjectFits(fstx.AddDefaultFits)
	app.Fit(mid.MyFitDemo)

	// 3. 根路由，中间件。
	// Note: 匹配到路由之后开始走这里的逻辑，执行过滤器
	app.Before(fstx.PmsParser)      // 解析请求参数，构造 ctx.Pms
	app.Before(jwtx.SdxSessBuilder) // 当前context中构造Session对象

	// +++++++++++++++++++++++++++++++++++++
	// 4.1 非登录组
	gpGhost := app.Group("/")
	ghost.AuthGroup(gpGhost)

	// +++++++++++++++++++++++++++++++++++++
	// 4.2 登录组。不同功能模块，分组对待
	gpAuth := app.Group("/")
	gpAuth.Before(jwtx.SdxMustLogin) // 验证当前请求是否已经登录

	auth.AdminGroup(gpAuth)
	auth.HRGroup(gpAuth)
	auth.CrmGroup(gpAuth)
}
