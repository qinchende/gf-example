package auth

import (
	"gf-example/web-demo/cf"
	"github.com/qinchende/gofast/cst"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/skill/httpx"
	"net/http"
	"time"
)

// curl -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/request_url
func RequestURL(c *fst.Context) {
	kv, err := httpx.DoRequestGetKVCtx(c.ReqRaw.Context(), &httpx.RequestPet{
		ProxyUrl: cf.Data.ProxyUrl,
		Method:   http.MethodGet,
		Url:      "http://127.0.0.1:8078/request_test_data",
		//QueryArgs: cst.KV{"tok": "t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"},
		//BodyArgs: cst.KV{"tok": "t:NDhDdjdwMEdaWTZoamtnY01o.RALE84mO4YGpAFdPfFEO8gi4NFcvH1kQV9IWmfaJuyc"},
	})
	time.Sleep(99 * time.Millisecond)
	c.FaiPanicIf(err != nil, err)
	c.SucData(kv["data"])
}

// curl -H "Content-Type: text/plain" -X GET http://127.0.0.1:8078/request_test_data
func RequestTestData(c *fst.Context) {
	time.Sleep(50 * time.Millisecond)
	c.SucData(cst.KV{"name": "GoFast"})
}
