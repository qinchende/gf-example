package auth

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func BeforeLogin(ctx *fst.Context) {
	logx.Info("Handler auth.BeforeLogin")
}

// curl -H "Content-Type: application/json" -X GET --data '{"name":"bmc","account":"rmb","age":37}' http://127.0.0.1:8078/login?account=admin\&pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data '{"name":"bmc","account":"rmb"}' http://127.0.0.1:8078/login?account=admin\&pass=abc123
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data "name=bmc&account=rmb&age=36" http://127.0.0.1:8078/login?account=admin\&pass=abc123
func LoginByAccPass(ctx *fst.Context) {
	// 模拟验证登录，写入 user_id
	account, _ := ctx.GetPms("account")
	pass, _ := ctx.GetPms("pass")

	if account == "admin" && pass == "abc" {
		ctx.Sess.Set("cus_id", 111)
		ctx.Sess.IsNew = false
		ctx.Sess.Save()
		ctx.SucKV(fst.KV{})
		return
	}
	ctx.FaiMsg("account and password error.")
}
