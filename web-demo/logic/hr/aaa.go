package hr

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)


func BeforeA(ctx *fst.Context) {
	logx.Info("Handler hr.BeforeA")
}
