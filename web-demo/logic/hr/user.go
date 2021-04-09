package hr

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func AddUser(ctx *fst.Context) {
	logx.Info("Handler hr.AddUser")

	ctx.SucKV(fst.KV{"name": "chen de"})
}

func AddDepartment(ctx *fst.Context) {
	logx.Info("Handler hr.AddDepartment")

	ctx.SucKV(fst.KV{"depart": "system"})
}
