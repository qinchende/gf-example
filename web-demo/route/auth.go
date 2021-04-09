package route

import (
	"gf-example/web-demo/logic/user"
	"github.com/qinchende/gofast/fst"
)

func noAuthGroup(gp *fst.RouterGroup) {
	gp.Before(user.BeforeA)
	gp.Get("/login", user.LoginDemo)
}
