package admin

import (
	"github.com/qinchende/gofast/cst"
	"github.com/qinchende/gofast/fst"
)

func BeforeA(c *fst.Context) {
	c.AddMsgBasket("Handler admin.BeforeA")
}

// curl -i -H "Content-Type: application/json" -X GET --data '{"name":"bmc","account":"rmb","age":37,"tok":"t:WUFZT3lKZFp5cmtlVkdQRnA2.UANo7oIqAyAw0P4Bwdzs2gcQumjgij2luC2jZrSLPOE"}' http://127.0.0.1:8078/admin/set
// curl -i -H "Content-Type: application/json" -X POST --data '{"name":"bmc","account":"rmb","age":37}' http://127.0.0.1:8078/admin/set
func SetParams(c *fst.Context) {
	c.AddMsgBasket("Handler admin.SetParams")
	c.SucData(cst.KV{"admin.SetParams": "suc"})
}
