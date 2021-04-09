package user

import "github.com/qinchende/gofast/fst"

func LoginDemo(ctx *fst.Context) {
	// 模拟登录，写入 user_id

	ctx.Sess.Set("cus_id", 111)
	ctx.Suc("{}")
}
