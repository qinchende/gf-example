package auth

import (
	"gf-example/web-demo/logic/admin"
	"github.com/qinchende/gofast/fst"
)

func AdminGroup(gp *fst.RouterGroup) {
	adm := gp.Group("/admin") // admin 组
	admA := adm.Group("/a")   // admin 组 下面又分 a 组
	admB := adm.Group("/b")   // admin 组 下面又分 b 组

	adm.Before(admin.BeforeA)
	adm.GetPost("/set", admin.SetParams) // GET 和 POST 同时支持

	admA.Get("/set", admin.SetParams)
	admB.Get("/set", admin.SetParams)
}
