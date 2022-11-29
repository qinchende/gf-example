package router

import (
	"gf-example/web-demo/logic/admin"
	"gf-example/web-demo/logic/auth"
	"gf-example/web-demo/logic/user"
	_ "gf-example/web-demo/model"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/sdx"
	"github.com/qinchende/gofast/sdx/mid"
)

func apiRoutes(app *fst.GoFast) {
	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 4.1 非登录组
	gpGhost := app.Group("/")
	//gpGhost.AfterMatch(user.AfterMatchRoute)

	// Get,Post支持单独定义配置参数
	get, post := gpGhost.GetPost("/mobile_code", auth.SendPhoneCode) // GET + POST 都支持
	get.Attrs(&mid.Attrs{TimeoutMS: 1000, MaxLen: 10240})            // 超时1秒，最大10K
	post.Attrs(&mid.Attrs{TimeoutMS: 600000})                        // 超时10分钟

	gpGhost.Get("/reg_by_email", user.RegByEmail)
	gpGhost.Post("/reg_by_mobile", user.RegByMobile)

	gpGhost.Post("/user/update/:user_id", user.UpdateBase).AfterMatch(user.AfterMatchRoute)     // 测试路由拦截
	gpGhost.Post("/user_update", user.UpdateBase)                                               // 更新
	gpGhost.Post("/query_users", user.QueryUser).B(user.BeforeQueryUser).A(user.AfterQueryUser) // 查询
	gpGhost.Get("/query_users", user.QueryUsers)
	gpGhost.Get("/query_users_cache", user.QueryUsersCache)
	gpGhost.GetPost("/query_user_gm", user.QueryGmInfo)

	gpGhost.Get("/bind_demo", user.BindDemo).B(user.BeforeBindDemo).A(user.AfterBindDemo).BeforeSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)
	gpGhost.Post("/bind_demo", user.BindDemo).B(user.BeforeBindDemo).A(user.AfterBindDemo).BeforeSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)

	gpGhost.Get("/request_test_data", auth.RequestTestData).Attrs(&mid.Attrs{TimeoutMS: 100})
	gpGhost.Get("/request_url", auth.RequestURL).Attrs(&mid.Attrs{TimeoutMS: 100})

	// 登录
	gpGhost.Get("/login", auth.LoginByAccPass).B(auth.BeforeLogin).Attrs(&mid.Attrs{TimeoutMS: 12000}) // 超时12秒
	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 4.2 登录组。不同功能模块，分组对待
	gpAuth := app.Group("/").B(sdx.SessMustLogin) // 检查当前请求是否已经登录
	gpAuth.Get("/logout", auth.Logout)            // logout

	// 4.3 Admin组 (也是需要先登录)
	adm := gpAuth.Group("/admin").B(admin.BeforeA) // admin 组
	adm.GetPost("/set", admin.SetParams)           // GET 和 POST 同时支持
}
