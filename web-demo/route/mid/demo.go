package mid

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"net/http"
)

// 自定义中间件函数
func MyFitDemo(w *fst.GFResponse, r *http.Request) {
	logx.Info("Middleware fit.my-demo.")
}
