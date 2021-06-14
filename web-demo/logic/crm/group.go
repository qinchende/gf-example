package crm

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func AddGroup(ctx *fst.Context) {
	logx.Info("Handler crm.AddGroup")

	ctx.SucKV(fst.KV{"group": "new"})
}
