package ghost

import (
	"gf-example/web-demo/logic/auth"
	"github.com/qinchende/gofast/fst"
)

// 这里的路由不需要登录
func AuthGroup(gp *fst.RouterGroup) {
	gp.Before(auth.BeforeLogin)
	gp.Get("/login/:id/:name", auth.LoginDemo)
}
