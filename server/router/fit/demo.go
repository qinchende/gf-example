package fit

import (
	"github.com/qinchende/gofast/aid/logx"
	"net/http"
)

// 自定义中间件函数

func MyFitDemo(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("HTTP middleware fit.my-demo.")
		next(w, r)
	}
}
