package router

import (
	"gf-example/server/logic"
	"gf-example/server/logic/admin"
	"gf-example/server/logic/auth"
	"gf-example/server/logic/nosess"
	"gf-example/server/logic/user"
	"gf-example/server/router/fit"
	"gf-example/server/router/web"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/sdx"
	"github.com/qinchende/gofast/sdx/mid"
)

func LoadRoutes(app *fst.GoFast) *fst.GoFast {
	// 1. 加载一些组合中间件
	app.Apply(web.StartEnd).Apply(sdx.SuperHandlers)

	// 2. 演示添加第一级中间件，第一级中间件不带fst.Context，而是原始的[req, res]
	app.UseHttpHandler(fit.RawHandlerDemo)

	// 3. 根路由，中间件。
	// Note: 匹配到路由之后开始走这里的逻辑，执行过滤器
	app.Before(sdx.PmsParser) // 解析请求参数，构造 ctx.Pms
	app.AfterMatch()          // 匹配路由之后，开始执行中间件链之前

	// 4. url api routes lists
	apiRoutes(app)

	return app
}

func apiRoutes(app *fst.GoFast) {
	app.BeforeSend(logic.AddHeaders)

	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 3.1 非session组，不理睬token信息
	gpNoToken := app.Group("/")
	gpNoToken.Get("/request_url", nosess.RequestURL).Bind(&mid.RAttrs{TimeoutMS: 100})
	gpNoToken.Get("/request_test_data", nosess.RequestTestData).Bind(&mid.RAttrs{TimeoutMS: 100})

	// 3.2 redis session 组
	gpToken := app.Group("/")
	gpToken.Before(sdx.TokSessBuilder) // 所有路由地址要携带tok信息，没有就自动分配一个

	// 3.2.1 非登录组，但有Session标记
	ghost := gpToken.Group("/")
	ghost.GetPost("/mobile_exist", user.MobileExist)
	ghost.GetPost("/email_exist", user.EmailExist)
	ghost.GetPost("/captcha_photo", user.CaptchaPhoto)
	ghost.GetPost("/mobile_valid_code", user.MobileValidCode)
	ghost.GetPost("/reg_by_mobile", user.RegByMobile)
	ghost.GetPost("/reg_by_email", user.RegByEmail)
	ghost.GetPost("/login1", auth.Login1)
	ghost.GetPost("/login2").Bind(auth.Login2())
	ghost.GetPost("/login3").Bind(auth.Login3())
	ghost.GetPost("/login4").Bind(auth.Login4())

	ghost.GetPost("/query_user_cache", user.QueryUserCache)
	ghost.GetPost("/query_users_cache", user.QueryUsersCache)

	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 3.2.2 登录组。不同功能模块，分组对待
	gpAuth := gpToken.Group("/")
	gpAuth.Before(sdx.SessMustLogin) // 所有路由地址需要先验证登录
	gpAuth.GetPost("/logout", auth.Logout)

	// 3.2.3 Admin组 (也是需要先登录)
	adm := gpAuth.Group("/admin")
	adm.B(admin.MustAdminPower)          // admin 组权限检查
	adm.GetPost("/set", admin.SetParams) // GET 和 POST 同时支持
}
