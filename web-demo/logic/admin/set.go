package admin

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func SetParams(ctx *fst.Context) {
	logx.Info("Handler admin.SetParams")

	ctx.SucKV(fst.KV{"set params": "suc"})
}
