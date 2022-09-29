package auth

import (
	"github.com/qinchende/gofast/cst"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/skill/httpx"
	"net/http"
)

// curl -H "Content-Type: application/json" -X GET --data '{"tok":"t:Q0JCM3R4dHhqWDZZM29FbTZr.xPEXaKSVK9nKwmhzOPIQzyqif1SnOhw68vTPj6024s"}' http://127.0.0.1:8078/request_url
func RequestURL(c *fst.Context) {
	kv, err := httpx.DoRequestGetKVCtx(c.ReqRaw.Context(), &httpx.RequestPet{
		Method:    http.MethodGet,
		Url:       "https://stk.tl50.com/get_kline_index_style?codes=SH603336",
		QueryArgs: cst.KV{"codes": "SH603338"},
		BodyArgs:  cst.KV{"codes": "SH603337"},
	})
	c.FaiPanicIf(err != nil, err)
	c.SucKV(map[string]any(kv))
	//c.SucStr("req url successful.")
}
