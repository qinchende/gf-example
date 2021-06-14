package crm

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func AddCustomer(ctx *fst.Context) {
	logx.Info("Handler crm.AddCustomer")

	ctx.SucKV(fst.KV{"customer": "zhang san"})
}
