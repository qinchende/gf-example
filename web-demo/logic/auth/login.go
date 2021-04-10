package auth

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func BeforeA(ctx *fst.Context) {
	logx.Info("Handler auth.BeforeA")
}

func LoginDemo(ctx *fst.Context) {
	// 模拟验证登录，写入 user_id
	account := ctx.Pms["account"]
	pass := ctx.Pms["pass"]

	if account == "admin" && pass == "abc" {
		ctx.Sess.Set("cus_id", 111)
		ctx.Suc("{}")
		return
	}
	ctx.Fai("account and password error.")
}
