package admin

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

// curl -H "Content-Type: application/json" -X GET --data '{"name":"bmc","account":"rmb","age":37}' http://127.0.0.1:8078/admin/set
// curl -H "Content-Type: application/json" -X POST --data '{"name":"bmc","account":"rmb","age":37}' http://127.0.0.1:8078/admin/set
func SetParams(ctx *fst.Context) {
	logx.Info("Handler admin.SetParams")

	ctx.SucKV(fst.KV{"set params": "suc"})
}
