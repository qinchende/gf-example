package auth

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
	"time"
)

// curl -i -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/mobile_code?len=6
// curl -i -H "Content-Type: application/x-www-form-urlencoded" -X POST --data "tok=t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s" http://127.0.0.1:8078/mobile_code?len=6
func SendPhoneCode(ctx *fst.Context) {
	// TODO: 1. 生成验证码 2. 调用短信通道发送
	kvs := fst.KV{"v_code": "123456"}
	ctx.Sess.SetKV(kvs)
	time.Sleep(100 * time.Millisecond)
	ctx.SucKV(kvs)
}

func BeforeLogin(c *fst.Context) {
	logx.Info("Handler auth.BeforeLogin")
}

// curl -H "Content-Type: application/json" -X GET --data '{"name":"bmc","account":"rmb","age":37,"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/login?account=admin\&pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data '{"name":"bmc","account":"rmb"}' http://127.0.0.1:8078/login?account=admin\&pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data "name=bmc&account=rmb&age=36" http://127.0.0.1:8078/login?account=admin\&pass=abc
func LoginByAccPass(c *fst.Context) {
	// 模拟验证登录，写入 user_id
	account := c.GetString("account")
	pass := c.GetString("pass")

	if account == "admin" && pass == "abc" {
		sdx.SessDestroy(c)
		sdx.SessRecreate(c)
		c.Sess.Set(sdx.MySess.GuidField, 111)
		_ = c.Sess.Save()
		c.SucKV(fst.KV{})
		return
	}
	c.FaiStr("account and password error.")
}
