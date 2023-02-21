package server

import (
	"gf-example/web-demo/model"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
)

func AppEnhance(app *fst.GoFast) {
	app.OnReady(func(fast *fst.GoFast) {
		logx.Info("App is ready.")
		model.InitModelsAttrs()
	})

	app.OnClose(func(fast *fst.GoFast) {
		logx.Info("App is closed.")
	})
}
