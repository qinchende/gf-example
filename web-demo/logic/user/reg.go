package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/store/sqlx"
)

// curl -H "Content-Type: application/json" -X POST --data '{"name":"陈德","account":"sdx","age":38,"v_code":"123456","email":"cd@qq.com","tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/reg_by_mobile
func RegByMobile(ctx *fst.Context) {
	// 通过自己判断字段合法性
	sVCode := ctx.Sess.Get("v_code")
	pVCode := ctx.Pms["v_code"]
	if sVCode == nil || sVCode == "" || pVCode == nil || pVCode == "" || sVCode != pVCode {
		ctx.FaiMsg("invalid mobile valid code")
		return
	}

	u := hr.SysUser{}
	if err := ctx.BindPms(&u); err != nil {
		ctx.FaiMsg(err.Error())
		return
	}
	logx.Info(u)

	//// 方式一：拼接sql语句。
	//// 注册，清理必要的数据，返回成功
	//r := config.MysqlZero.Exec("insert into sys_user(account,name,age,nickname,created_at,updated_at)values(?, ?, ?, ?, now(), now())",
	//	u.Account, u.Name, u.Age, u.Nickname)
	//id, _ := r.LastInsertId()

	// 方式二：Gorm 三方包保存
	//ret := cf.GormZero.Create(&u)
	//if ret.Error != nil {
	//	ctx.FaiMsg("Created err: " + ret.Error.Error())
	//	return
	//}
	//u.Age = 49
	//cf.GormZero.Updates(&u)

	//gormUsers := make([]hr.SysUserDemo, 0)
	//cf.GormZero.Find(&gormUsers, "age=91")
	//logx.Info(gormUsers)
	//
	//gormUsers2 := new([]*hr.SysUserDemo)
	//cf.GormZero.Find(gormUsers2, "age=91")
	//logx.Info(gormUsers2)

	//ctx.SucKV(fst.KV{"id": u.ID, "affected": ret.RowsAffected})
	//return

	// 方式三：GoFast自带ORM功能
	ct := cf.Zero.Insert(&u)
	if ct > 0 {
		logx.Infof("Insert success, new id: %d", u.ID)
	}

	u.Name = "chende"
	ct = cf.Zero.Update(&u)

	u.Name = "wang"
	u.Age = 78
	u.Status = 1
	ct = cf.Zero.UpdateColumns(&u, "name", "status")
	ct = cf.Zero.UpdateColumns(&u, "age,status")

	newUser := hr.SysUser{}
	ct = cf.Zero.QueryID(&newUser, u.ID)
	logx.Info(newUser)

	ct = cf.Zero.QueryRow(&newUser, "id=?", u.ID)
	logx.Info(newUser)

	myUsers := make([]*hr.SysUser, 0)
	ct = cf.Zero.QueryRows(&myUsers, "age=? and status=?", 91, 1)
	logx.Info(myUsers[0])

	myUsers2 := new([]*hr.SysUser)
	ct = cf.Zero.QueryRows(myUsers2, "age=? and status=?", 38, 0)
	logx.Info((*myUsers2)[0])
	ct = cf.Zero.QueryRows2(myUsers2, "age,name", "age=78 and status=0 limit 5")
	logx.Info((*myUsers2)[0])

	records := new([]fst.KV)
	ct = cf.Zero.QueryPet(records, &sqlx.SelectPet{
		//Sql: "select * from sys_user where age=? and status=0",
		Table:   "sys_user",
		Columns: "id,name,age,status",
		Offset:  1,
		Limit:   9,
		Where:   "age=? and status=0",
		Prams:   []interface{}{78},
	})
	if ct > 0 {
		logx.Info((*records)[0])
	}

	cf.Zero.QueryPetCC(records, &sqlx.SelectPetCC{
		CacheType: sqlx.CacheMem,
		SelectPet: sqlx.SelectPet{
			Table:   "sys_user",
			Columns: "id,name,age,status",
			Offset:  1,
			Limit:   9,
			Where:   "age=? and status=0",
			Prams:   []interface{}{78},
		},
	})
	if len(*records) > 0 {
		logx.Info((*records)[0])
	}

	ct = cf.Zero.Delete(&u)
	ctx.SucKV(fst.KV{"id": u.ID, "updated_at": u.UpdatedAt, "R": (*records)[0]})
	return
}

// curl -H "Content-Type: application/json" -X POST --data '{"name":"陈德","account":"sdx","age":38,"v_code":"123456","email":"cd@qq.com","tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/reg_by_email
func RegByEmail(ctx *fst.Context) {
	sVCode := ctx.Sess.Get("v_code")
	pVCode := ctx.Pms["v_code"]
	if sVCode == nil || sVCode == "" || pVCode == nil || pVCode == "" || sVCode != pVCode {
		ctx.FaiMsg("invalid mobile valid code")
		return
	}

	u := hr.SysUser{}
	if err := ctx.BindPms(&u); err != nil {
		ctx.FaiMsg(err.Error())
		return
	}
	logx.Info(u)

	// 数据库事务的测试
	//zero := cf.Zero.TransBegin()
	//defer zero.TransEnd()
	//zero.Insert(&u)
	//myUsers := make([]*hr.SysUser, 0)
	//ct := cf.Zero.QueryRows(&myUsers, "age=? and status=?", 91, 1)
	//logx.Info(ct)

	myUsers := make([]*hr.SysUser, 0)
	cf.Zero.TransFunc(func(zero *sqlx.MysqlORM) {
		zero.Insert(&u)
		logx.Info(u)

		ct := cf.Zero.QueryRows(&myUsers, "age=? and status=?", 91, 1)
		logx.Info(ct)
	})

	ctx.SucKV(fst.KV{"record": *myUsers[0]})
	return
}
