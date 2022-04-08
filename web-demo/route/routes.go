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
	"github.com/qinchende/gofast/jwtx"
)

func routesList(app *fst.GoFast) {
	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 4.1 非登录组
	gpGhost := app.Group("/")
	gpGhost.Get("/login", auth.LoginByAccPass).Before(auth.BeforeLogin).Config(&mid.RConfig{Timeout: 12000})

	// Get,Post支持单独定义配置参数
	get, post := gpGhost.GetPost("/mobile_code", sms.SendPhoneCode)
	get.Config(&mid.RConfig{Timeout: 1000, MaxLen: 10000})
	post.Config(&mid.RConfig{Timeout: 321000})

	gpGhost.Post("/reg_by_mobile", user.RegByMobile)
	gpGhost.Post("/reg_by_email", user.RegByEmail)
	gpGhost.Get("/user_list", hr.UserList)

	//// GET
	//gpGhost.Get("/bind_demo", user.BindDemo).Before(user.BeforeBindDemo).After(user.AfterBindDemo).PreSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)
	//// POST
	//gpGhost.Post("/bind_demo", user.BindDemo).Before(user.BeforeBindDemo).After(user.AfterBindDemo).PreSend(user.BeforeBindDemoSend).AfterSend(user.AfterBindDemoSend)

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
}
