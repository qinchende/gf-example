package route

import (
	"gf-example/web-demo/logic/auth"
	"github.com/qinchende/gofast/fst"
)

// 这里的路由不需要登录
func noAuthGroup(gp *fst.RouterGroup) {
	gp.Before(auth.BeforeA)
	gp.Get("/login", auth.LoginDemo)
}
