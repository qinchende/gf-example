package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func BeforeA(ctx *fst.Context) {
	logx.Info("Handler user.BeforeA")
}

// curl -H "Content-Type: application/json" -X POST --data '' http://127.0.0.1:8078/query_user
func QueryUser(ctx *fst.Context) {
	ccUser := hr.SysUser{}
	ct := cf.Zero.QueryIDCC(&ccUser, 150)
	logx.Info(ct)
	logx.Info(ccUser)

	ctx.SucKV(fst.KV{"record": ccUser})
	return
}
