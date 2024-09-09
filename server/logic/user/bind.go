package user

import (
	"github.com/qinchende/gofast/aid/logx"
	"github.com/qinchende/gofast/fst"
)

func BeforeBindDemo(c *fst.Context) {
	logx.Info("ghost.auth.BeforeBindDemo")
}

func AfterBindDemo(c *fst.Context) {
	logx.Info("ghost.auth.AfterBindDemo")
}

func BeforeBindDemoSend(c *fst.Context) {
	logx.Info("ghost.auth.BeforeBindDemoSend")
}

func AfterBindDemoSend(c *fst.Context) {
	logx.Info("ghost.auth.AfterBindDemoSend")
}

// curl -H "Content-Type: application/json" -d '{"name":"bmc","account":"rmb","age":38}' http://127.0.0.1:8078/bind_demo?first=chen\&last=de\&tok=t:THNqNjVFTU5sbkNtd0N3OXRp.6UWKmsqPhnrGAbOk7zeRtsUW0uhptj4gI5/FiiIylAs
// curl -H "Content-Type: application/json" -d '{"name":"bmc","account":"rmb","age":37}' http://127.0.0.1:8078/bind_demo?ids[a]=1234\&ids[b]=hello\&first=chen\&last=de
// curl -H "Content-Type: application/x-www-form-urlencoded" -d '{"name":"bmc","account":"rmb"}' http://127.0.0.1:8078/bind_demo?first=chen\&last=de
// curl -H "Content-Type: application/x-www-form-urlencoded" -d "name=bmc&account=rmb&age=36" http://127.0.0.1:8078/bind_demo?ids[a]=1234\&ids[b]=hello\&first=chen\&last=de
//func BindDemo(c *fst.Context) {
//	user := &hr.SysUser{}
//	c.PanicIfErr(c.Bind(&user), nil)
//
//	title := hr.Title{}
//	//fst.GFPanicErr(c.Bind(&title))
//	c.PanicIfErr(c.Bind(&title), cf.FaiBindError)
//
//	//// query url 中的参数
//	//ids := c.QueryMap("ids")
//	//first := c.QueryDef("first", "Guest")
//	//last := c.Query("last")
//	//
//	//// post form 提交的参数
//	//acc := c.PostForm("account")
//	//name := c.PostFormDef("name", "anonymous")
//	//age := c.PostFormMap("age")
//
//	c.SucData(cst.KV{
//		"uname":      user.Name,
//		"nickname":   user.Nickname,
//		"account":    user.Account,
//		"age":        user.Age,
//		"title_name": title.Name,
//		//"q_ids":      ids,
//		//"q_first":    first,
//		//"q_last":     last,
//		//"f_acc":      acc,
//		//"f_name":     name,
//		//"f_age":      age,
//	})
//}
