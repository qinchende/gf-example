package auth

import (
	"github.com/qinchende/gofast/fst"
)

// curl -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/logout
func Logout(ctx *fst.Context) {
	ctx.DestroySession()
	ctx.NewSession()
	ctx.SucStr("logout success. ")
}
