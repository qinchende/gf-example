package user

import (
	"gf-example/web-demo/cf"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/cst"
	"github.com/qinchende/gofast/fst"
	"net/http"
	"regexp"
	"strings"
)

// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s","user_name":"陈德11"}' http://127.0.0.1:8078/user/update/11
// 进一步检查匹配到的特定路由，而先不走可能无用的中间件
func AfterMatchRoute(c *fst.Context) {
	if c.Req.Referer() != "https://www.tl50.com/" {
		c.AbortDirect(http.StatusHTTPVersionNotSupported, "只支持http/1.x协议")
		// c.SetRouteTo404()
		return
	} else if !strings.HasPrefix(c.Req.RemoteAddr, "10.10") {
		c.SetRouteTo404()
		return
	} else if c.Req.ProtoMajor == 2 {
		c.AbortDirect(http.StatusHTTPVersionNotSupported, "只支持http/1.x协议")
		return
	}

	uid := c.UrlParam("user_id")
	if regexp.MustCompile("^1[0-9]$").MatchString(uid) == false {
		c.AbortDirect(http.StatusNotFound, "路由匹配失败")
		return
	}
}

// curl -H "Content-Type: application/json" -X POST --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s","user_name":"陈德11","user_id":"11"}' http://127.0.0.1:8078/user_update
func UpdateBase(c *fst.Context) {
	userId := c.GetIntMust("user_id")
	u := hr.SysUser{}
	ct := cf.Zero.QueryPrimaryCache(&u, userId)
	cst.PanicIf(ct <= 0, cf.FaiNotFound)

	newName := c.GetStringMust("user_name")
	u.Name = newName
	if ct = cf.Zero.UpdateFields(&u, "Name", "Status"); ct <= 0 {
		c.FaiRet(cf.FaiUserUpdate)
		//c.FaiCode(rt.FaiUserUpdateError)
		//c.FaiMsg("更新失败")
	} else {
		c.SucData(cst.KV{"id": u.ID, "name": u.Name})
	}
}
