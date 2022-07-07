package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s","user_name":"闪电","user_id":"11"}' http://127.0.0.1:8078/user_update
func UpdateBase(ctx *fst.Context) {
	userId := ctx.Pms["user_id"]
	newName := ctx.Pms["user_name"]

	ccUser := hr.SysUser{}
	ct := cf.Zero.QueryIDCC(&ccUser, userId)

	ccUser.Name = newName.(string)
	cf.Zero.UpdateColumns(&ccUser, "name")

	logx.Info(ct)
	logx.Info(ccUser)

	ctx.SucKV(fst.KV{"id": ccUser.ID, "name": ccUser.Name})
	return
}

// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/query_users
func QueryUser(ctx *fst.Context) {
	ccUser := hr.SysUser{}
	ct := cf.Zero.QueryIDCC(&ccUser, 11)
	logx.Info(ct)
	logx.Info(ccUser)

	ctx.SucKV(fst.KV{"id": ccUser.ID, "name": ccUser.Name})
	return
}
