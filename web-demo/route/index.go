package route

import (
	"gf-example/web-demo/logic/admin"
	"gf-example/web-demo/logic/auth"
	"gf-example/web-demo/logic/crm"
	"gf-example/web-demo/logic/hr"
	"gf-example/web-demo/logic/user"
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

	// 2. 全局中间件（拦截器）
	// Note: 请求进来，并没有定位到具体的路由。就需要走这些过滤器
	// 所有的请求都要走这里指定的拦截器，发生错误就直接中断返回
	app.Fits(fstx.AddDefaultFits) // 默认的一组中间件
	app.Fit(mid.MyFitDemo)        // 自定义中间件

	// 3. 根路由，中间件。
	// Note: 匹配到路由之后开始走这里的逻辑，执行过滤器
	app.Before(fstx.PmsParser)      // 解析请求参数，构造 ctx.Pms
	app.Before(jwtx.SdxSessBuilder) // 当前context中构造Session对象

	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 4.1 非登录组
	gpGhost := app.Group("/")
	gpGhost.Before(auth.BeforeLogin)
	gpGhost.Get("/login", auth.LoginByAccPass)

	// GET
	gpGhost.Get("/bind_demo", user.BindDemo).Before(user.BeforeBindDemo).After(user.AfterBindDemo).PreSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)
	// POST
	gpGhost.Post("/bind_demo", user.BindDemo).Before(user.BeforeBindDemo).After(user.AfterBindDemo).PreSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)

	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 4.2 登录组。不同功能模块，分组对待
	gpAuth := app.Group("/")
	gpAuth.Before(jwtx.SdxMustLogin) // 检查当前请求是否已经登录

	// logout
	gpAuth.Get("/logout", auth.Logout)

	// Admin
	adm := gpAuth.Group("/admin").Before(admin.BeforeA) // admin 组
	adm.GetPost("/set", admin.SetParams)                // GET 和 POST 同时支持
	admA := adm.Group("/a")                             // admin 组 下面又分 a 组
	admB := adm.Group("/b")                             // admin 组 下面又分 b 组
	admA.Get("/set", admin.SetParams)
	admB.Get("/set", admin.SetParams)

	// HR
	hrGroup := gpAuth.Group("/hr")
	hrGroup.Before(hr.BeforeA)
	//hrGroup.Get("/add_user", hr.AddUser)
	//hrGroup.Get("/add_depart", hr.AddDepartment)

	// CRM
	crmGroup := gpAuth.Group("/crm")
	crmGroup.Before(crm.BeforeA)
	//crmGroup.Get("/add_user", crm.AddCustomer)
	//crmGroup.Get("/add_depart", crm.AddGroup)
}
