package auth

import (
	"gf-example/server/cf"
	"gf-example/server/model/acc"
	"github.com/qinchende/gofast/aid/logx"
	"github.com/qinchende/gofast/core/cst"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/sdx"
	"github.com/qinchende/gofast/store/gson"
	"strconv"
	"time"
)

// curl -i -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/mobile_code?len=6
// curl -i -H "Content-Type: application/x-www-form-urlencoded" -X POST --data "tok=t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s" http://127.0.0.1:8078/mobile_code?len=6
func SendPhoneCode(c *fst.Context) {
	// TODO: 1. 生成验证码 2. 调用短信通道发送
	kvs := cst.WebKV{"v_code": "123456"}
	c.Sess.SetValues(kvs)
	time.Sleep(100 * time.Millisecond)
	c.SucData(kvs)
}

func BeforeLogin(c *fst.Context) {
	logx.Info("Handler auth.BeforeLogin")
}

// curl -H "Content-Type: application/json" -X GET --data '{"name":"bmc","account":"rmb","age":37,"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/login?account=admin\&pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data '{"name":"bmc","account":"rmb"}' http://127.0.0.1:8078/login?account=admin\&pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data "name=bmc&account=rmb&age=36" http://127.0.0.1:8078/login?account=admin\&pass=abc
func LoginByAccPass(c *fst.Context) {
	// 模拟验证登录，写入 user_id
	account := c.GetStringMust("account")
	pass := c.GetStringMust("pass")

	if account == "admin" && pass == "abc" {
		c.Sess.Destroy()
		c.Sess.Recreate()
		c.Sess.SetUid("111")
		c.Sess.Save()
		c.SucData(cst.KV{})
		return
	}
	c.FaiMsg("account and password error.")
}

// curl -H "Content-Type: application/json" -X POST -d '{"name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw","name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw"}' http://127.0.0.1:8019/login1?account=admin\&hash_pass=abc
// curl -H "Content-Type: application/json" -X POST -d '{"name":"bmc","account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw"}' http://127.0.0.1:8019/login1?account=admin\&hash_pass=abc
// curl -H "Content-Type: application/json" -X POST -d '{"name":"bmc","account":"rmb","age":37,"tok":"eyJleHAiOiIxNzA2NTY0NTAzIiwidWlkIjoiMSIsImp0aSI6ImRERmlWbU4xWTJwTWQydzNPSEpzY2pVeCJ9.rh0q40Z0rO3Ze_Qdm4fEaQ"}' http://127.0.0.1:8019/login1?account=admin\&hash_pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST -d '{"name":"bmc","account":"rmb"}' http://127.0.0.1:8019/login1?account=admin\&hash_pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST -d "name=bmc&account=rmb&age=36" http://127.0.0.1:8019/login1?account=admin\&hash_pass=abc
func Login1(c *fst.Context) {
	type params struct {
		//Account  string `v:"must,len=[1:64],match=email"`
		Account  string `v:"must,len=[1:64]"`
		HashPass string `v:"must,len=[1:255]"`
	}

	var in params
	c.PanicIfErr(c.BindAndValid(&in), "数据解析错误")

	var uid int64
	cf.DDemo.QuerySqlRow(&uid, acc.SqlAuthLogin, in.Account, in.HashPass)

	if uid > 0 {
		ss := c.Sess.(*sdx.TokSession)
		ss.Destroy()
		ss.Recreate()
		ss.SetUid(strconv.FormatInt(uid, 10))
		ss.Save()
		c.SucRet(cf.SucLogin)
		return
	}
	c.FaiRet(cf.FaiLoginAccPassErr)
}

// curl -H "Content-Type: application/json" -X POST -d '{"name":{"first":"chen","second":"de"},"account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw"}' http://127.0.0.1:8019/login2?account=admin\&hash_pass=abc
// curl -H "Content-Type: application/json" -X POST -d '{"pages":[11,22,33,44,55],"account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw"}' http://127.0.0.1:8019/login2?account=admin\&hash_pass=abc
func Login2() *fst.RHandler {
	type params struct {
		//Pages    []int8
		Name struct {
			First  string `v:"must"`
			Second string `v:"must"`
		}
		Account  string `v:"must,len=[1:64]"`
		HashPass string `v:"must,len=[1:255]"`
	}

	handler := func(c *fst.Context) {
		var in params
		c.PanicIfErr(c.BindAndValid(&in), "数据解析错误")

		var uid int64
		cf.DDemo.QuerySqlRow(&uid, acc.SqlAuthLogin, in.Account, in.HashPass)

		if uid > 0 {
			c.Sess.Destroy()
			c.Sess.Recreate()
			c.Sess.SetUid(strconv.FormatInt(uid, 10))
			c.Sess.Save()
			c.SucRet(cf.SucLogin)
			return
		}
		c.FaiRet(cf.FaiLoginAccPassErr)
	}

	return fst.WrapHandler(handler, nil, sdx.PmsKeys(&params{}))
}

// curl -H "Content-Type: application/json" -X POST -d '{"name":{"first":"chen","second":"de"},"account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw"}' http://127.0.0.1:8019/login3?account=admin\&hash_pass=abc
// curl -H "Content-Type: application/json" -X POST -d '{"pages":[11,22,33,44,55],"account":"rmb","age":37,"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw"}' http://127.0.0.1:8019/login3?account=admin\&hash_pass=abc
func Login3() *fst.RHandler {
	type params struct {
		sdx.BaseFields
		//Pages []int8
		Name struct {
			First  string `v:"must"`
			Second string `v:"must"`
		}
		Account  string `v:"must,len=[1:64]"`
		HashPass string `v:"must,len=[1:255]"`
	}

	handler := func(c *fst.Context) {
		in, err := fst.PmsAsAndValid[params](c)
		c.PanicIfErr(err, "数据解析错误")

		var uid int64
		cf.DDemo.QuerySqlRow(&uid, acc.SqlAuthLogin, in.Account, in.HashPass)

		if uid > 0 {
			c.Sess.Destroy()
			c.Sess.Recreate()
			c.Sess.SetUid(strconv.FormatInt(uid, 10))
			c.Sess.Save()
			c.SucRet(cf.SucLogin)
			return
		}
		c.FaiRet(cf.FaiLoginAccPassErr)
	}

	return fst.WrapHandler(handler, fst.NewSuperKV[params], nil)
}

// curl -H "Content-Type: application/json" -X POST -d '{"gson":[1,1,["name","account","age","login","mobile","tok"],[["b{m}c","bmcrmb",37,true,"1344466338783","t:Q0J44CM3R4dHhqWDZZM2944FbTZr"]]],"tok":"c2llUXp1T21hUUNyUVNoVTcw.3cvfzRHmKJN8CMjfrQqcJw"}' http://127.0.0.1:8019/login4?account=admin\&hash_pass=abc
func Login4() *fst.RHandler {
	type params struct {
		sdx.BaseFields
		Account  string `v:"must,len=[1:1]"`
		HashPass string `v:"must,len=[1:1]"`
		Gson     gson.RowsDecPet
	}

	handler := func(c *fst.Context) {
		in := fst.PmsAs[params](c)

		var uid int64
		cf.DDemo.QuerySqlRow(&uid, acc.SqlAuthLogin, in.Account, in.HashPass)

		if uid > 0 {
			c.Sess.Destroy()
			c.Sess.Recreate()
			c.Sess.SetUid(strconv.FormatInt(uid, 10))
			c.Sess.Save()
			c.SucRet(cf.SucLogin)
			return
		}
		c.FaiRet(cf.FaiLoginAccPassErr)
	}

	return fst.WrapHandler(handler, func() cst.SuperKV {
		in := &params{}
		in.Gson.List = new([]acc.SysUser)
		return fst.ToSuperKV(in)
	}, nil)
}
