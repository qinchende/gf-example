package sms

import (
	"github.com/qinchende/gofast/fst"
	"time"
)

// curl -i -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/mobile_code?len=6
// curl -i -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/mobile_code?len=6
// curl -i -H "Content-Type: application/x-www-form-urlencoded" -X POST --data "tok=t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s" http://127.0.0.1:8078/mobile_code?len=6
func SendPhoneCode(ctx *fst.Context) {
	// TODO: 1. 生成验证码 2. 调用短信通道发送
	kvs := fst.KV{"v_code": "123456"}
	ctx.Sess.SetKV(kvs)
	time.Sleep(100 * time.Millisecond)
	ctx.SucKV(kvs)
}
