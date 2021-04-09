package route

import (
	"gf-example/web-demo/logic/crm"
	"github.com/qinchende/gofast/fst"
)

func crmGroup(gp *fst.RouterGroup) {
	crmGroup := gp.Group("/crm")

	crmGroup.Before(crm.BeforeA)
	crmGroup.Get("/add_user", crm.AddCustomer)
	crmGroup.Get("/add_depart", crm.AddGroup)
}
