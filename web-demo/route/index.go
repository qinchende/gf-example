package route

import (
	"gf-example/web-demo/route/mid"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/fstx"
	"github.com/qinchende/gofast/jwtx"
)

func LoadRoutes(app *fst.GoFast) {
	// 1. 基础
	serverSetup(app)

	// 2. 全局中间件（拦截器）
	// 请求进来，并没有定位到具体的路由。就需要走这些过滤器
	app.RegFits(fstx.AddDefaultFits)
	app.Fit(mid.FitDemo)

	// 3. 根路由，中间件。
	// 匹配到路由之后开始走这里的逻辑，执行过滤器
	app.Before(jwtx.SdxSessBuilder) // 还原当前请求的Session对象

	// 4.1 非登录组
	groupA := app.Group("/")
	noAuthGroup(groupA)

	// 4.2 登录组。不同功能模块，分组对待
	groupB := app.Group("/")
	groupB.Before(jwtx.SdxMustLogin) // 验证当前请求是否已经登录

	hrGroup(groupB)
	crmGroup(groupB)
}
