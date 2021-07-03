package auth

import (
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func BeforeLogin(ctx *fst.Context) {
	logx.Info("Handler auth.BeforeLogin")
}

func LoginDemo(ctx *fst.Context) {
	// 模拟验证登录，写入 user_id
	account := ctx.Pms["account"]
	pass := ctx.Pms["pass"]

	if account == "admin" && pass == "abc" {
		ctx.Sess.Set("cus_id", 111)
		ctx.Suc("{}")
		return
	}
	ctx.Fai("account and password error.")
}

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
// curl -H "Content-Type: application/json" -X POST --data '{"name":"bmc","account":"rmb","age":37}' http://127.0.0.1:8078/bind_demo?first=chen\&last=de
// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST --data '{"name":"bmc","account":"rmb"}' http://127.0.0.1:8078/bind_demo?first=chen\&last=de
// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST --data "name=bmc&account=rmb&age=36" http://127.0.0.1:8078/bind_demo?ids[a]=1234\&ids[b]=hello\&first=chen\&last=de
func BindDemo(ctx *fst.Context) {
	user := hr.User{}
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
	firstname := ctx.DefaultQuery("first", "Guest")
	lastname := ctx.Query("last") // shortcut for ctx.Request.URL.Query().Get("lastname")

	//message := ctx.PostForm("account")
	//nick := ctx.DefaultPostForm("name", "anonymous")
	//names := ctx.PostFormMap("age")

	ctx.Suc(fst.KV{
		"uname":      user.Name,
		"nickname":   user.Nickname,
		"account":    user.Account,
		"age":        user.Age,
		"title_name": title.Name,
		"ids":        ids,
		"firstname":  firstname,
		"lastname":   lastname,
	})
	//ctx.String(http.StatusOK, fmt.Sprintf("file uploaded!"))
	//ctx.JSON(http.StatusOK, myData)
}
