package user

import (
	"gf-example/server/cf"
	"github.com/qinchende/gofast/core/dts"
	"github.com/qinchende/gofast/fst"
)

//type DemoUser struct {
//	Name    string `json:"name"`
//	Title   string `json:"title"`
//	Contact struct {
//		Home string `json:"home"`
//		Cell string `json:"cell"`
//	} `json:"contact"`
//}
//
//var jsonStr = `{"name": "Gopher","title": "Programmer", "contact": {"home": "415.333.3333", "cell": "415.555.5555"}}`
//
//func TestJsonDecode(c *fst.Context) {
//	user := DemoUser{}
//	_ = bind.BindJsonBytes(&user, lang.STB(jsonStr), bind.AsReq)
//
//	c.SucMsg("Just TestJsonDecode.")
//}

// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET --data "name=bmc&account=rmb&age=36" http://127.0.0.1:8019/mobile_exist?mobile=13466663333
func MobileExist(c *fst.Context) {
	mobile := c.GetStringMust("mobile")
	c.PanicIf(!dts.IsMobile(mobile), "请输入正确手机号")

	ct := cf.DDemo.QuerySqlInt64("select count(id) from sys_user where mobile=? limit 1;", mobile)
	c.IfSucFai(ct >= 1, "手机号已注册", "手机号未注册")
}

func EmailExist(c *fst.Context) {
	email := c.GetStringMust("email")
	c.PanicIf(!dts.IsEmail(email), "请输入正确的电子邮箱")

	ct := cf.DDemo.QuerySqlInt64("select count(id) from sys_user where email=? limit 1;", email)
	c.IfSucFai(ct >= 1, "邮箱已注册", "邮箱未注册")
}
