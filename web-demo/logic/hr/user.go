package hr

import (
	"gf-example/web-demo/config"
	"gf-example/web-demo/model/hr"
	"github.com/qinchende/gofast/fst"
)

//func AddUser(c *fst.Context) {
//	logx.Info("Handler hr.AddUser")
//
//	newUser := hr.User{}
//	if err := c.BindPms(&newUser); err != nil {
//		c.FaiMsg("参数错误")
//		return
//	}
//
//	newTitle := hr.Title{}
//	if err := c.BindPms(&newTitle); err != nil {
//		c.FaiMsg("参数错误")
//		return
//	}
//
//	newDepart := hr.Department{}
//	if err := c.BindPms(&newDepart); err != nil {
//		c.FaiMsg("参数错误")
//		return
//	}
//
//	c.Suc(0, "Saved.", newUser)
//	//c.SucKV(fst.KV{"name": "chen de"})
//}

// curl -H "Content-Type: application/json" -X GET http://127.0.0.1:8078/user_list
func UserList(c *fst.Context) {
	var users []hr.User
	config.GormZero.Find(&users, []int{1, 2})

	c.Suc(0, "two", users)
}
