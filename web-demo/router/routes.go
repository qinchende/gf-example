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

	// Get,Post支持单独定义配置参数
	get, post := gpGhost.GetPost("/mobile_code", auth.SendPhoneCode) // GET + POST 都支持
	get.Config(&mid.RConfig{Timeout: 1000, MaxLen: 10240})           // 超时1秒，最大10K
	post.Config(&mid.RConfig{Timeout: 600000})                       // 超时10分钟

	gpGhost.Post("/reg_by_mobile", user.RegByMobile)
	gpGhost.Get("/reg_by_email", user.RegByEmail)
	gpGhost.Post("/reg_by_email", user.RegByEmail)

	gpGhost.Post("/user_update", user.UpdateBase)                                                        // 更新
	gpGhost.Post("/query_users", user.QueryUser).Before(user.BeforeQueryUser).After(user.AfterQueryUser) // 查询
	gpGhost.Get("/query_users", user.QueryUsers)
	gpGhost.Get("/query_users_cache", user.QueryUsersCache)
	gpGhost.GetPost("/query_user_gm", user.QueryGmInfo)

	gpGhost.Get("/login", auth.LoginByAccPass).Before(auth.BeforeLogin).Config(&mid.RConfig{Timeout: 12000}) // 超时12秒

	// GET
	gpGhost.Get("/bind_demo", user.BindDemo).Before(user.BeforeBindDemo).After(user.AfterBindDemo).PreSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)
	// POST
	gpGhost.Post("/bind_demo", user.BindDemo).Before(user.BeforeBindDemo).After(user.AfterBindDemo).PreSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)

	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 4.2 登录组。不同功能模块，分组对待
	gpAuth := app.Group("/").Before(sdx.SessMustLogin) // 检查当前请求是否已经登录
	gpAuth.Get("/logout", auth.Logout)                 // logout

	// 4.3 Admin组 (也是需要先登录)
	adm := gpAuth.Group("/admin").Before(admin.BeforeA) // admin 组
	adm.GetPost("/set", admin.SetParams)                // GET 和 POST 同时支持
}
