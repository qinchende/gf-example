package user

import (
	"gf-example/server/cf"
	"github.com/mojocn/base64Captcha"
	"github.com/qinchende/gofast/aid/randx"
	"github.com/qinchende/gofast/core/cst"
	"github.com/qinchende/gofast/core/dts"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/fst/httpx"
	"net/http"
)

// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET http://127.0.0.1:8019/captcha_photo
func CaptchaPhoto(c *fst.Context) {
	var store = base64Captcha.DefaultMemStore
	var driver base64Captcha.Driver
	driver = &base64Captcha.DriverString{
		Length: 6,
		Width:  200,
		Height: 100,
		Source: "abcdefghijklmnopqrstABCDEFG",
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := captcha.Generate()
	c.PanicIfErr(err, "验证码图片生成是吧")

	c.SucData(cst.KV{"data": b64s, "captchaId": id})
}

// curl -H "Content-Type: application/x-www-form-urlencoded" -X GET http://127.0.0.1:8019/mobile_valid_code?mobile=13466663333
func MobileValidCode(c *fst.Context) {
	mobile := c.GetStringMust("mobile")
	c.PanicIf(!dts.IsMobile(mobile), "请正确输入手机号")

	// Note: 1. 生成验证码 2. 调用短信通道发送
	vCode := randx.RandomNumbers(6)
	kvs := cst.WebKV{"v_mobile": mobile, "v_code": vCode}

	// 发送管理模块，发送通知
	retKV, err := httpx.DoRequestGetKVCtx(c.Req.Raw.Context(), &httpx.RequestPet{
		ProxyUrl:  cf.DParams.ProxyUrl,
		Method:    http.MethodGet,
		Url:       cf.DParams.MmsSendUrl,
		QueryArgs: kvs,
		//BodyArgs: cst.KV{"tok": "t:NDhDdjdwMEdaWTZoamtnY01o.RALE84mO4YGpAFdPfFEO8gi4NFcvH1kQV9IWmfaJuyc"},
	})
	c.PanicIfErr(err, "发送申请失败")

	if retKV["status"] == "suc" {
		c.Sess.SetValues(kvs)
		c.Sess.Save()
		c.SucMsg("验证码发送成功")
	} else {
		c.FaiMsg("验证码发送失败")
	}
}
