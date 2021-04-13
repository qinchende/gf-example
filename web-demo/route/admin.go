package route

import (
	"gf-example/web-demo/logic/admin"
	"github.com/qinchende/gofast/fst"
)

func adminGroup(gp *fst.RouterGroup) {
	admGroup := gp.Group("/admin") // admin 组
	admA := admGroup.Group("/a")   // admin 组 下面又分 a 组
	admB := admGroup.Group("/b")   // admin 组 下面又分 b 组

	admGroup.Before(admin.BeforeA)
	admGroup.GetPost("/set", admin.SetParams) // GET 和 POST 同时支持

	admA.Get("/set", admin.SetParams)
	admB.Get("/set", admin.SetParams)
}
