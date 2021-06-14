package hr

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func AddDepartment(ctx *fst.Context) {
	logx.Info("Handler hr.AddDepartment")

	ctx.SucKV(fst.KV{"depart": "system"})
}
