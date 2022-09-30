package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/cf/rt"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
)

// 测试自定义表的主键
// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s","user_id":"11"}' http://127.0.0.1:8078/query_user_gm
func QueryGmInfo(c *fst.Context) {
	userId := c.GetIntMust("user_id")
	userGm := hr.SysUserGmInfo{}
	ct := cf.Zero.QueryPrimaryCache(&userGm, userId)

	userGm.IsOpen = 0
	cf.Zero.Update(&userGm)

	//newGm := hr.SysUserGmInfo{UserID: 12, IsOpen: 2, OpenTime: lang.Ptr(time.Now())}
	//yn := cf.Zero.Insert(&newGm)
	//c.FaiPanicIf(yn <= 0, "add error")

	c.FaiPanicIf(ct <= 0, rt.FaiNotFound)
	c.SucData(fst.KV{"result": userGm})
}
