package user

import (
	"gf-example/web-demo/config"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
)

// curl -H "Content-Type: application/json" -X POST --data '{"name":"rmb","account":"bmc","age":37,"v_code":"123456","tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/reg_by_mobile
func RegByMobile(ctx *fst.Context) {
	// 通过自己判断字段合法性
	sVCode := ctx.Sess.Get("v_code")
	pVCode := ctx.Pms["v_code"]
	if sVCode == nil || sVCode == "" || pVCode == nil || pVCode == "" || sVCode != pVCode {
		ctx.FaiMsg("invalid mobile valid code")
		return
	}

	u := hr.User{}
	if err := ctx.BindPms(&u); err != nil {
		ctx.FaiMsg("BindPms err: " + err.Error())
		return
	}

	//// 方式一：拼接sql语句。
	//// 注册，清理必要的数据，返回成功
	//r := config.MysqlZero.Exec("insert into sys_users(account,name,age,nickname,created_at,updated_at)values(?, ?, ?, ?, now(), now())",
	//	u.Account, u.Name, u.Age, u.Nickname)
	//id, _ := r.LastInsertId()
	//ctx.SucKV(fst.KV{"id": id})
	//return

	// 方式二：ORM保存
	ret := config.GormZero.Create(&u)
	if ret.Error != nil {
		ctx.FaiMsg("Created err: " + ret.Error.Error())
		return
	}
	ctx.SucKV(fst.KV{"id": u.ID, "affected": ret.RowsAffected})
	return
}
