package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/cst"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/store/sqlx"
)

// curl -H "Content-Type: application/json" -d '{"name":"陈德","account":"sdx","age":38,"v_code":"123456","email":"cd@qq.com","tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/reg_by_mobile
func RegByMobile(c *fst.Context) {
	// 通过自己判断字段合法性
	sVCode := c.Sess.Get("v_code")
	pVCode := c.GetMust("v_code")
	if sVCode == nil || sVCode == "" || pVCode == nil || pVCode == "" || sVCode != pVCode {
		c.FaiMsg("invalid mobile valid code")
		return
	}

	u := hr.SysUser{}
	cst.PanicIfErr(c.Bind(&u))
	logx.Infos(u)

	//// 方式一：拼接sql语句。
	//// 注册，清理必要的数据，返回成功
	//r := config.MysqlZero.Exec("insert into sys_user(account,name,age,nickname,created_at,updated_at)values(?, ?, ?, ?, now(), now())",
	//	u.Account, u.Name, u.Age, u.Nickname)
	//id, _ := r.LastInsertId()

	// 方式二：Gorm 三方包保存
	//ret := cf.GormZero.Create(&u)
	//if ret.Error != nil {
	//	c.FaiMsg("Created err: " + ret.Error.Error())
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

	//c.SucData(fst.KV{"id": u.ID, "affected": ret.RowsAffected})
	//return

	// 方式三：GoFast自带ORM功能
	ct := cf.Zero.Insert(&u)
	if ct > 0 {
		logx.InfoF("Insert success, new id: %d", u.ID)
	}

	u.Name = "chende"
	ct = cf.Zero.Update(&u)

	u.Name = "wang"
	u.Age = 78
	u.Status = 1
	ct = cf.Zero.UpdateFields(&u, "Name", "Status")
	ct = cf.Zero.UpdateFields(&u, "Age,Status")

	newUser := hr.SysUser{}
	ct = cf.Zero.QueryPrimary(&newUser, u.ID)
	logx.Infos(newUser)

	ct = cf.Zero.QueryRow(&newUser, "id=?", u.ID)
	logx.Infos(newUser)

	myUsers := make([]*hr.SysUser, 0)
	ct = cf.Zero.QueryRows(&myUsers, "age=? and status=?", 91, 1)
	if len(myUsers) > 0 {
		logx.Infos(myUsers[0])
	}

	myUsers2 := new([]*hr.SysUser)
	ct = cf.Zero.QueryRows(myUsers2, "age=? and status=?", 38, 0)
	if len(*myUsers2) > 0 {
		logx.Infos((*myUsers2)[0])
	}
	ct = cf.Zero.QueryRows2(myUsers2, "age,name", "age=78 and status=0 limit 5")
	if len(*myUsers2) > 0 {
		logx.Infos((*myUsers2)[0])
	}

	records := new([]cst.KV)
	ct = cf.Zero.QueryPet(&sqlx.SelectPet{
		List: records,
		//Sql: "select * from sys_user where age=? and status=0",
		Table:   "sys_user",
		Columns: "id,name,age,status",
		Offset:  1,
		Limit:   9,
		Where:   "age=? and status=0",
		Args:    []any{78},
	})
	if ct > 0 {
		logx.Infos((*records)[0])
	}

	ct = cf.Zero.Delete(&u)
	c.SucData(cst.KV{"id": u.ID, "updated_at": u.UpdatedAt})
	return
}

// curl -H "Content-Type: application/json" -X GET --data '{"name":"陈德","account":"sdx","age":38,"v_code":"123456","email":"cd@qq.com","tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/reg_by_email?ids=abc\&ids=123
func RegByEmail(c *fst.Context) {
	sVCode := c.Sess.Get("v_code")
	pVCode := c.GetMust("v_code")
	if sVCode == nil || sVCode == "" || pVCode == nil || pVCode == "" || sVCode != pVCode {
		c.FaiMsg("invalid mobile valid code")
		return
	}

	u := hr.SysUser{}
	cst.PanicIfErr(c.Bind(&u))
	logx.Infos(u)

	// 第一种事务
	zero := cf.Zero.TransBegin()
	defer zero.TransEnd()
	zero.Insert(&u)
	myUsers := make([]*hr.SysUser, 0)
	ct := zero.QueryRows(&myUsers, "age=? and status=?", 38, 3)
	logx.Infos(ct)

	// 第二种事务
	//myUsers := make([]*hr.SysUser, 0)
	//cf.Zero.TransFunc(func(zero *sqlx.MysqlORM) {
	//	zero.Insert(&u)
	//	logx.Info(u)
	//	ct := zero.QueryRows(&myUsers, "age=? and status=?", 38, 3)
	//	logx.Info(ct)
	//})

	cst.PanicIf(ct <= 0, "无记录")
	c.SucData(cst.KV{"record": *myUsers[0]})
}
