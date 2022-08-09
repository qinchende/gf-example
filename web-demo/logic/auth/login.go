package auth

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"github.com/qinchende/gofast/sdx"
)

func BeforeLogin(c *fst.Context) {
	logx.Info("Handler auth.BeforeLogin")
}

// curl -H "Content-Type: application/json" -X GET --data '{"name":"bmc","account":"rmb","age":37,"tok":"t:QnBQTHNDT3RIS2V2aFJyUk1o.rEnZy6QeaS/fDtG3Kj/eBBwKDRbfJs8/nAqIxtmzdM"}' http://127.0.0.1:8078/login?account=admin\&pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data '{"name":"bmc","account":"rmb"}' http://127.0.0.1:8078/login?account=admin\&pass=abc
// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data "name=bmc&account=rmb&age=36" http://127.0.0.1:8078/login?account=admin\&pass=abc
func LoginByAccPass(c *fst.Context) {
	// 模拟验证登录，写入 user_id
	account := c.GetString("account")
	pass := c.GetString("pass")

	if account == "admin" && pass == "abc" {
		sdx.DestroySession(c)
		sdx.NewSession(c)
		c.Sess.Set(sdx.MySess.GuidField, 111)
		c.Sess.Save()
		c.SucKV(fst.KV{})
		return
	}
	c.FaiStr("account and password error.")
}
