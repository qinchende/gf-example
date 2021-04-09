package route

import (
	"gf-example/web-demo/logic/hr"
	"github.com/qinchende/gofast/fst"
)

func hrGroup(gp *fst.RouterGroup) {
	hrGroup := gp.Group("/hr")

	hrGroup.Before(hr.BeforeA)
	hrGroup.Get("/add_user", hr.AddUser)
	hrGroup.Get("/add_depart", hr.AddDepartment)
}
