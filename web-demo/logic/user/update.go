package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/cf/rd"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
)

// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s","user_name":"陈德11","user_id":"11"}' http://127.0.0.1:8078/user_update
func UpdateBase(c *fst.Context) {
	userId := c.GetIntMust("user_id")
	u := hr.SysUser{}
	ct := cf.Zero.QueryPrimaryCache(&u, userId)
	c.FaiPanicIf(ct <= 0, rd.FaiNotFound)

	newName := c.GetStringMust("user_name")
	u.Name = newName
	if ct = cf.Zero.UpdateFields(&u, "Name", "Status"); ct <= 0 {
		c.FaiCode(rd.FaiUserUpdateError)
		//c.FaiStr("更新失败")
	} else {
		c.SucKV(fst.KV{"id": u.ID, "name": u.Name})
	}
}
