package admin

import (
	"gf-example/server/cf"
	"github.com/qinchende/gofast/core/cst"
	"github.com/qinchende/gofast/fst"
)

func BeforeA(c *fst.Context) {
	c.CarryMsg("Handler admin.BeforeA")
}

func MustAdminPower(c *fst.Context) {
	c.FaiRet(cf.FaiNeedAdminPower)
}

// curl -i -H "Content-Type: application/json" -X GET --data '{"name":"bmc","account":"rmb","age":37,"tok":"t:WUFZT3lKZFp5cmtlVkdQRnA2.UANo7oIqAyAw0P4Bwdzs2gcQumjgij2luC2jZrSLPOE"}' http://127.0.0.1:8019/admin/set
// curl -i -H "Content-Type: application/json" -X POST --data '{"name":"bmc","account":"rmb","age":37}' http://127.0.0.1:8019/admin/set
func SetParams(c *fst.Context) {
	c.CarryMsg("Handler admin.SetParams")
	c.SucData(cst.KV{"admin.SetParams": "suc"})
}
