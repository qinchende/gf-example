package route

import (
	"gf-example/web-demo/logic/admin"
	"github.com/qinchende/gofast/fst"
)

func adminGroup(gp *fst.RouterGroup) {
	admGroup := gp.Group("/admin")
	admA := admGroup.Group("/a")
	admB := admGroup.Group("/b")

	admGroup.Before(admin.BeforeA)
	admGroup.GetPost("/set", admin.SetParams) // GET 和 POST 同时支持

	admA.Get("/set", admin.SetParams)
	admB.Get("/set", admin.SetParams)
}
