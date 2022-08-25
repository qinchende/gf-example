package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
)

// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s","user_name":"陈德12","user_id":"12"}' http://127.0.0.1:8078/user_update
func UpdateBase(c *fst.Context) {
	userId := c.MustGet("user_id").(string)
	newName := c.MustGet("user_name").(string)

	ccUser := hr.SysUser{}
	cf.Zero.QueryIDCache(&ccUser, userId)

	ccUser.Name = newName
	cf.Zero.UpdateColumns(&ccUser, "name")

	//logx.Info(ct)
	//logx.Info(ccUser)

	c.SucKV(fst.KV{"id": ccUser.ID, "name": ccUser.Name})
}
