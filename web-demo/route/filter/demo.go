package filter

import (
	"github.com/qinchende/gofast/logx"
	"net/http"
)

// 自定义中间件函数

func MyFitDemo(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("Middleware fit.my-demo.")
		next(w, r)
	}
}
