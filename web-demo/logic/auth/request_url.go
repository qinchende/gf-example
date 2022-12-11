package auth

import (
	"github.com/qinchende/gofast/cst"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/skill/httpx"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

// curl -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/request_url
func RequestURL(c *fst.Context) {
	kv, err := httpx.DoRequestGetKVCtx(c.ReqRaw.Context(), &httpx.RequestPet{
		//ProxyUrl:  cf.Data.ProxyUrl,
		Method:    http.MethodGet,
		Url:       "http://127.0.0.1:8078/request_test_data",
		QueryArgs: cst.KV{"tok": "t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"},
		//BodyArgs: cst.KV{"tok": "t:NDhDdjdwMEdaWTZoamtnY01o.RALE84mO4YGpAFdPfFEO8gi4NFcvH1kQV9IWmfaJuyc"},
	})
	time.Sleep(99 * time.Millisecond)
	c.FaiPanicIf(err != nil, err)
	c.SucData(kv["data"])
}

var randTool = rand.New(rand.NewSource(time.Now().UnixNano()))
var count int32 = 0

// curl -H "Content-Type: text/plain" -X GET http://127.0.0.1:8078/request_test_data
// curl -H "Content-Type: text/plain" -X GET http://127.0.0.1:8078/request_test_data?tok=t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s
func RequestTestData(c *fst.Context) {
	ct := atomic.AddInt32(&count, 1)

	var timeout int32 = 1
	if ct < 33 {
		timeout += ct * 3
	} else if ct < 400 {
		timeout += (50 + randTool.Int31n(200))
	} else if ct < 500 {
		timeout += (500 - ct)
	} else {
		timeout = 10
	}
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	c.SucData(cst.KV{"Count": ct, "Timeout": timeout})
}
