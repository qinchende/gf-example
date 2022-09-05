package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/skill/timex"
	"github.com/qinchende/gofast/store/sqlx"
	"time"
)

func BeforeQueryUser(c *fst.Context) {
	return
	// c.FaiStr("error: before QueryUser")
	// c.AbortFaiStr("error: before abort")

	// 这里测试一下 sqlx 的预处理连接
	userTest := hr.SysUser{}
	sqlStr := "select * from sys_user where id=?;"

	startTime := timex.Now()
	myStmt := cf.Zero.Prepare(sqlStr, true)
	for i := 11; i <= 12; i++ {
		ct := myStmt.QueryRow(&userTest, i)
		if ct <= 0 {
			logx.InfoF("User id: %#v can't find.", i)
			continue
		}
		logx.InfoF("User id: %#v exist. Name is %s", i, userTest.Name)
	}
	myStmt.Close()
	dur := timex.Since(startTime)
	logx.InfoF("[SQL Prepare][%dms]", dur/time.Millisecond)
}

// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/query_users
// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s","user_id":"12"}' http://127.0.0.1:8078/query_users
func QueryUser(c *fst.Context) {
	userId := c.GetIntMust("user_id")

	ccUser := hr.SysUser{}
	ct := cf.Zero.QueryIDCache(&ccUser, userId)
	c.AddMsgBasket("The info will show in log ext section.")

	if ct > 0 {
		c.SucKV(fst.KV{"id": ccUser.ID, "name": ccUser.Name})
	} else {
		c.FaiStr("can't find the record")
	}
}

func AfterQueryUser(c *fst.Context) {
	return
	// c.FaiStr("error: after QueryUser")

	// 这里测试一下 sqlx 的非预处理方案
	userTest := hr.SysUser{}
	sqlStr := "select * from sys_user where id=?;"

	startTime := timex.Now()
	for i := 11; i <= 12; i++ {
		sqlRows := cf.Zero.QuerySql(sqlStr, i)
		ct := sqlx.ScanRow(&userTest, sqlRows)
		sqlx.ErrLog(sqlRows.Close())

		if ct <= 0 {
			logx.InfoF("User id: %#v can't find.", i)
			continue
		}
		logx.InfoF("User id: %#v exist. Name is %s", i, userTest.Name)
	}
	dur := timex.Since(startTime)
	logx.InfoF("[SQL No Prepare][%dms]", dur/time.Millisecond)
}

// curl -H "Content-Type: application/json" -X GET --data '{"name":"bmc"}' http://127.0.0.1:8078/query_users
func QueryUsers(c *fst.Context) {
	myUsers := make([]*hr.SysUser, 0)
	ct := cf.Zero.QueryPet(&sqlx.SelectPet{
		Target: &myUsers,
		//Sql: "select * from sys_user where age=? and status=0",
		//Table:   "sys_user",
		Columns: "*",
		Where:   "age=? and status=? and id=?",
		Args:    []any{38, 3, 11},
		Limit:   500,
	})
	logx.Infos(ct)

	ct2 := cf.Zero.QueryPet(&sqlx.SelectPet{
		Target: &myUsers,
		Sql:    "select id,name,age,status from sys_user where age=? and status=? and id=?",
		Args:   []any{38, 3, 11},
	})
	logx.Infos(ct2)

	curCt, totalCt := cf.Zero.QueryPetPaging(&sqlx.SelectPet{
		Target:   &myUsers,
		Columns:  "id,name,age,status",
		Where:    "age=?",
		Args:     []any{38},
		Page:     1,
		PageSize: 10,
		OrderBy:  "id desc",
		GroupBy:  "id",
		PetCache: &sqlx.PetCache{
			ExpireS:   3600,
			CacheType: sqlx.CacheRedis,
		},
	})
	logx.Infos(curCt, ",", totalCt)

	c.SucKV(fst.KV{"records": myUsers})
}

// curl -H "Content-Type: application/json" -X GET --data '{"name":"bmc"}' http://127.0.0.1:8078/query_users_cache
func QueryUsersCache(c *fst.Context) {
	myUsers := make([]*hr.SysUser, 0)
	// 自动缓存结果集
	myPet := &sqlx.SelectPet{
		Target: &myUsers,
		Sql:    "select * from sys_user where age=? and status=?",
		Args:   []any{38, 3},
		PetCache: &sqlx.PetCache{
			ExpireS:   3600,
			CacheType: sqlx.CacheRedis,
		},
	}
	ct3 := cf.Zero.QueryPet(myPet)
	//_ = cf.Zero.DeletePetCache(myPet)
	logx.Infos(ct3)

	c.SucKV(fst.KV{"records": myUsers})
}
