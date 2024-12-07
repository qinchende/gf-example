package main

import (
	"fmt"
	"github.com/qinchende/gofast/aid/conf"
	"github.com/qinchende/gofast/aid/logx"
	"github.com/qinchende/gofast/core/cst"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/sdx"
	"log"
	"net/http"
	"time"
)

func MainSample() {
	var handler = func(str string) func(c *fst.Context) {
		return func(c *fst.Context) {
			log.Println(str)
		}
	}

	var handlerRender = func(str string) func(c *fst.Context) {
		return func(c *fst.Context) {
			log.Println(str)
			c.FaiData(cst.KV{"data": str})
		}
	}

	appCfg := &fst.ServerConfig{
		RunMode: "debug",
	}
	_ = conf.LoadFromJson(&appCfg.WebConfig, []byte("{}"))
	_ = conf.LoadFromJson(&appCfg.LogConfig, []byte("{}"))
	appCfg.WebConfig.PrintRouteTrees = true
	appCfg.LogConfig.LogMedium = "console"
	appCfg.LogConfig.LogLevel = "debug"
	appCfg.LogConfig.LogStyle = "sdx-json"
	appCfg.LogConfig.FilePath = "_logs_"

	app := fst.CreateServer(appCfg)
	logx.SetupDefault(&appCfg.LogConfig)

	// 拦截器，微服务治理 ++++++++++++++++++++++++++++++++++++++
	app.UseHttpHandler(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("app enter before 1")
			next(w, r)
			log.Println("app enter after 1")
		}
	})
	app.Apply(sdx.SuperHandlers)
	app.Before(sdx.PmsParser) // 解析请求参数，构造 ctx.Pms

	// 根路由
	app.Reg404(func(c *fst.Context) {
		c.Json(http.StatusNotFound, "404-Can't find the path.")
	})
	app.Reg405(func(c *fst.Context) {
		c.Json(http.StatusMethodNotAllowed, "405-Method not allowed.")
	})

	// ++ (用这种方法可以模拟中间件需要上下文变量的场景)
	app.Before(func(c *fst.Context) {
		c.Set("nowTime", time.Now())
		//time.Sleep(3 * time.Second)
	})
	app.After(func(c *fst.Context) {
		// 处理后获取消耗时间
		val, exist := c.Get("nowTime")
		if exist {
			costTime := time.Since(val.(time.Time)) / time.Millisecond
			fmt.Printf("The request cost %dms", costTime)
		}
	})
	// ++ end

	// curl -H "Content-Type: application/json" -X POST  --data '{"data":"bmc","nick":"yes"}' http://127.0.0.1:8099/root?first=yang\&last=lmx
	// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST  --data "data=bmc&nick=yes" http://127.0.0.1:8099/root?
	// 		ids[a]=1234\&ids[b]=hello\&first=yang\&last=lmx
	type MyData struct {
		Data string `v:"must,len=[1:16]"`
		Nick string `v:"must,len=[1:32]"`
	}
	app.Post("/root", func(c *fst.Context) {
		var myData MyData
		c.PanicIfErr(c.BindAndValid(&myData), "数据解析错误")
		log.Printf("%v %+v %#v\n", myData, myData, myData)

		ids, _ := c.GetString("ids")
		firstname := c.GetStringDef("first", "Guest")
		lastname := c.GetStringMust("last")

		message := c.GetStringMust("data")
		nick := c.GetStringDef("nick", "anonymous")

		//names := c.PostFormMap("names")
		c.SucData(cst.KV{
			"message": message,
			"nick":    nick,
			"first":   firstname,
			"last":    lastname,
			"ids":     ids,
			//"data":    myData,
		})
		//c.String(http.StatusOK, fmt.Sprintf("file uploaded!"))
		//c.JSON(http.StatusOK, myData)
	})

	//app.Post("/root", handler("root"))
	app.Before(handler("before root")).After(handler("after root"))

	// 分组路由1
	adm := app.Group("/admin")
	adm.After(handler("after group admin")).Before(handler("before group admin"))

	tst := adm.Get("/sdx", handlerRender("handle sdx"))
	// 添加路由处理事件
	tst.Before(handler("before tst_url"))
	tst.After(handler("after tst_url"))
	tst.BeforeSend(handler("beforeSend tst_url"))
	tst.AfterSend(handler("afterSend tst_url"))

	// 分组路由2
	adm2 := app.Group("/admin2").Before(handler("before admin2"))
	adm2.Get("/zht", handler("zht")).After(handler("after zht"))

	adm22 := adm2.Group("/group2").Before(handler("before group2"))
	adm22.Get("/lmx", handler("lmx")).Before(handler("before lmx"))

	// 应用级事件
	app.OnReady(func(fast *fst.GoFast) {
		log.Println("App OnReady Call.")
		log.Printf("Listening and serving HTTP on %s\n", "127.0.0.1:8099")
	})
	app.OnClose(func(fast *fst.GoFast) {
		log.Println("App OnClose Call.")
	})
	// 开始监听接收请求
	app.Listen("127.0.0.1:8099")
}
