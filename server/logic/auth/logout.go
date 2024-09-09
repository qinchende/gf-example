package auth

import (
	"gf-example/server/cf"
	"github.com/qinchende/gofast/fst"
)

// curl -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/logout
func Logout(c *fst.Context) {
	c.PanicPet = cf.FaiLogoutPanic // 异常处理方法一
	c.Sess.Destroy()               // 这个肯定成功，否则底层就抛异常了
	c.Sess.Recreate()              // 新的token
	c.SucRet(cf.SucLogout)         // 退出系统
}
