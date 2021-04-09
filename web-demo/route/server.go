package route

import (
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/logx"
	"net/http"
)

func serverSetup(app *fst.GoFast) {
	// 应用级事件
	app.OnReady(func(fast *fst.GoFast) {
		logx.Info("App OnReady Call.")
	})
	app.OnClose(func(fast *fst.GoFast) {
		logx.Info("App OnClose Call.")
	})

	// 根路由 +++++++++++++++++++++++++++++++++++++++++++++++
	app.NoRoute(func(ctx *fst.Context) {
		ctx.JSON(http.StatusNotFound, "404-Can't find the path.")
	})
	app.NoMethod(func(ctx *fst.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, "405-Method not allowed.")
	})
}
