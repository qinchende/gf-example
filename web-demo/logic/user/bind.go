package user

import (
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func BeforeBindDemo(ctx *fst.Context) {
	logx.Info("ghost.auth.BeforeBindDemo")
}

func AfterBindDemo(ctx *fst.Context) {
	logx.Info("ghost.auth.AfterBindDemo")
}

func BeforeBindDemoSend(ctx *fst.Context) {
	logx.Info("ghost.auth.BeforeBindDemoSend")
}

func AfterBindDemoSend(ctx *fst.Context) {
	logx.Info("ghost.auth.AfterBindDemoSend")
}

// curl -H "Content-Type: application/json" -X POST --data '{"name":"bmc","account":"rmb","age":38}' http://127.0.0.1:8078/bind_demo?first=chen\&last=de\&tok=t:THNqNjVFTU5sbkNtd0N3OXRp.6UWKmsqPhnrGAbOk7zeRtsUW0uhptj4gI5/FiiIylAs
// curl -H "Content-Type: application/json" -X POST --data '{"name":"bmc","account":"rmb","age":37}' http://127.0.0.1:8078/bind_demo?ids[a]=1234\&ids[b]=hello\&first=chen\&last=de
// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST --data '{"name":"bmc","account":"rmb"}' http://127.0.0.1:8078/bind_demo?first=chen\&last=de
// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST --data "name=bmc&account=rmb&age=36" http://127.0.0.1:8078/bind_demo?ids[a]=1234\&ids[b]=hello\&first=chen\&last=de
func BindDemo(ctx *fst.Context) {
	user := hr.SysUser{}
	if err := ctx.BindPms(&user); err != nil {
		ctx.FaiErr(err)
		return
	}
	logx.Infof("%v %+v %#v\n", user, user, user)

	var title hr.Title
	if err := ctx.BindPms(&title); err != nil {
		ctx.FaiErr(err)
		return
	}

	ids := ctx.QueryMap("ids")
	first := ctx.DefaultQuery("first", "Guest")
	last := ctx.Query("last") // shortcut for ctx.Request.URL.Query().Get("lastname")

	acc := ctx.PostForm("account")
	name := ctx.DefaultPostForm("name", "anonymous")
	age := ctx.PostFormMap("age")

	ctx.SucKV(fst.KV{
		"uname":      user.Name,
		"nickname":   user.Nickname,
		"account":    user.Account,
		"age":        user.Age,
		"title_name": title.Name,
		"q_ids":      ids,
		"q_first":    first,
		"q_last":     last,
		"f_acc":      acc,
		"f_name":     name,
		"f_age":      age,
	})
}
