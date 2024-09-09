package user

import (
	"gf-example/server/cf"
	"gf-example/server/model/acc"
	"github.com/qinchende/gofast/aid/hashx"
	"github.com/qinchende/gofast/aid/lang"
	"github.com/qinchende/gofast/fst"
)

// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET http://127.0.0.1:8019/reset_pass?account=admin\&new_hash_pass=abc
func ResetPass(c *fst.Context) {
	newHashPass := c.GetStringMust("new_hash_pass")

	uPass := acc.SysUserPass{}
	ct := cf.DDemo.QueryPrimary(&uPass, c.Sess.GetUid())
	c.PanicIf(ct <= 0, "找不到用户")

	salt := lang.BTS(hashx.Md5(lang.STB(lang.ToString(uPass.ID + int64(110)))))
	uPass.Sha1Pass = lang.BTS(hashx.Sha1(lang.STB(salt + newHashPass)))
	ct = cf.DDemo.UpdateColumns(&uPass, "hash_pass")
	c.IfSucFai(ct > 0, "密码修改成功", "新密码保存失败")
}
