package route

import (
	"gf-example/web-demo/logic/admin"
	"gf-example/web-demo/logic/auth"
	"gf-example/web-demo/logic/crm"
	"gf-example/web-demo/logic/hr"
	"gf-example/web-demo/logic/sms"
	"gf-example/web-demo/logic/user"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/fst/mid"
	"github.com/qinchende/gofast/sdx/jwtx"
)

func apiRoutes(app *fst.GoFast) {
	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 4.1 非登录组
	gpGhost := app.Group("/")

	// Get,Post支持单独定义配置参数
	get, post := gpGhost.GetPost("/mobile_code", sms.SendPhoneCode) // GET + POST 都支持
	get.Config(&mid.RConfig{Timeout: 1000, MaxLen: 10240})          // 超时1秒，最大10K
	post.Config(&mid.RConfig{Timeout: 600000})                      // 超时10分钟

	gpGhost.Post("/reg_by_mobile", user.RegByMobile)
	gpGhost.Get("/reg_by_email", user.RegByEmail)
	gpGhost.Post("/reg_by_email", user.RegByEmail)

	gpGhost.Post("/user_update", user.UpdateBase)                                                        // 更新
	gpGhost.Post("/query_users", user.QueryUser).Before(user.BeforeQueryUser).After(user.AfterQueryUser) // 查询

	gpGhost.Get("/login", auth.LoginByAccPass).Before(auth.BeforeLogin).Config(&mid.RConfig{Timeout: 12000}) // 超时12秒

	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 4.2 登录组。不同功能模块，分组对待
	gpAuth := app.Group("/")
	gpAuth.Before(jwtx.SdxMustLogin) // 检查当前请求是否已经登录

	// logout
	gpAuth.Get("/logout", auth.Logout)

	// Admin
	adm := gpAuth.Group("/admin").Before(admin.BeforeA) // admin 组
	adm.GetPost("/set", admin.SetParams)                // GET 和 POST 同时支持
	//admA := adm.Group("/a")                             // admin 组 下面又分 a 组
	//admB := adm.Group("/b")                             // admin 组 下面又分 b 组
	//admA.Get("/set", admin.SetParams)
	//admB.Get("/set", admin.SetParams)

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

	//// GET
	//gpGhost.Get("/bind_demo", user.BindDemo).Before(user.BeforeBindDemo).After(user.AfterBindDemo).PreSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)
	//// POST
	//gpGhost.Post("/bind_demo", user.BindDemo).Before(user.BeforeBindDemo).After(user.AfterBindDemo).PreSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)

}
