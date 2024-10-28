package web

import (
	"gf-example/server/model"
	"github.com/qinchende/gofast/aid/logx"
	"github.com/qinchende/gofast/fst"
)

func StartEnd(app *fst.GoFast) *fst.GoFast {
	app.OnReady(func(fast *fst.GoFast) {
		logx.Info("App is ready.")
		model.InitModelsAttrs()
		//time.Sleep(3 * time.Second)
		//runtime.GC() // 主动垃圾回收，看是否有不安全内存操作
	})

	app.OnClose(func(fast *fst.GoFast) {
		logx.Info("App is closed.")
	})

	return app
}
