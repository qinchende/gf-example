package ghost

import (
	"gf-example/web-demo/logic/auth"
	"github.com/qinchende/gofast/fst"
)

// 这里的路由不需要登录
func AuthGroup(gp *fst.RouterGroup) {
	gp.Before(auth.BeforeLogin)
	gp.Get("/login/:id/:name", auth.LoginDemo)

	// POST
	gp.Post("/bind_demo", auth.BindDemo).Before(auth.BeforeBindDemo).After(auth.AfterBindDemo).PreSend(auth.BeforeBindDemoSend).AfterSend(auth.AfterBindDemoSend)
	// GET
	bindGet := gp.Get("/bind_demo", auth.BindDemo)
	bindGet.Before(auth.BeforeBindDemo).After(auth.AfterBindDemo).PreSend(auth.BeforeBindDemoSend).AfterSend(auth.AfterBindDemoSend)
}
