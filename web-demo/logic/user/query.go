package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
)

func BeforeQueryUser(c *fst.Context) {
	//c.FaiStr("error: before QueryUser")
	//c.AbortFaiStr("error: before abort")
}

// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/query_users
// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s","user_id":"12"}' http://127.0.0.1:8078/query_users
func QueryUser(c *fst.Context) {
	userId := c.MustGet("user_id").(string)

	ccUser := hr.SysUser{}
	ct := cf.Zero.QueryIDCC(&ccUser, userId)

	c.AddMsgBasket("The info will show in log ext section.")

	if ct > 0 {
		c.SucKV(fst.KV{"id": ccUser.ID, "name": ccUser.Name})
	} else {
		c.FaiStr("can't find the record")
	}
}

func AfterQueryUser(c *fst.Context) {
	//c.FaiStr("error: after QueryUser")
}
