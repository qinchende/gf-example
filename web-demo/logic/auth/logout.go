package auth

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/sdx"
)

// curl -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/logout
func Logout(c *fst.Context) {
	sdx.SessDestroy(c)
	c.SucStr("logout success. ")
}
