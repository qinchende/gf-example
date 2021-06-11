package auth

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func BeforeLogin(ctx *fst.Context) {
	logx.Info("Handler auth.BeforeLogin")
}

type item struct {
	Key   string
	Value string
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
