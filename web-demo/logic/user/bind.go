package user

import (
	"fmt"
	"github.com/gin-gonic/gin/render"
	"github.com/qinchende/gofast/fst"
	"github.com/qinchende/gofast/jwtx"
	"github.com/qinchende/gofast/logx"
	"log"
)

func ReqBindingDemos(app *fst.GoFast) {
	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// 新建一个分组
	bind := app.Group("/")
	// 本分组需要登录验证。
	bind.Before(jwtx.SdxMustLogin)
	// 分组事件
	bind.PreSend(func(ctx *fst.Context) {
		kv := (*ctx.PRender).(render.JSON).Data.(fst.KV)
		kv["others"] = "cd.net by yes"

		logx.Info("pre send")
	})
	bind.AfterSend(func(ctx *fst.Context) {
		//kv := (*ctx.PRender).(render.JSON).Data.(fst.KV)
		//logx.Info(kv)
		logx.Info("after send")
	})
	bind.PreBind(func(ctx *fst.Context) {
		logx.Info("pre bind")
	})

	// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	// curl -i http://127.0.0.1:8099/root?first=chen\&last=de
	// 标准路由写法
	rtItem := bind.Get("/root2", func(c *fst.Context) {
		// panic("this root is broken")
		firstname := c.Query("first")
		lastname := c.Query("last") // shortcut for c.Request.URL.Query().Get("lastname")

		c.Sess.Set("first", firstname)
		c.Sess.Set("last", lastname)
		c.Sess.Set("age", 19.88)
		c.Sess.Delete("age")

		//time.Sleep(5 * time.Second)

		logx.Info("handler")
		c.Suc(fst.KV{
			"first": firstname,
			"last":  lastname,
		})
	})
	rtItem.PreSend(func(ctx *fst.Context) {
		logx.Info("pre send")
	})
	rtItem.AfterSend(func(ctx *fst.Context) {
		logx.Info("after send")
	})

	// curl -H "Content-Type: application/json" -X POST --data '{"data":"bmc","nick":"yes"}' http://127.0.0.1:8099/root?first=yang\&last=lmx\&tok=t:V0tLc2RWQWo0VGVWSkpwVXM0.KUSVSsRfXo7Kzucc8D3DYn7iXyMYAl0WZwzhAWI5fQ
	// curl -H "Content-Type: application/json" -X POST --data '{"data":"bmc","nick":"yes"}' http://127.0.0.1:8099/root?first=yang\&last=lmx
	// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST --data '{"data":"bmc","nick":"yes"}' http://127.0.0.1:8099/root?first=yang\&last=lmx
	// curl -H "Content-Type: application/x-www-form-urlencoded" -X POST --data "data=bmc&nick=yes" http://127.0.0.1:8099/root?ids[a]=1234\&ids[b]=hello\&first=yang\&last=lmx
	type MyData struct {
		First string `json:"first" binding:"required"`
		Last  string `json:"last"`
	}
	bind.Get("/root", func(ctx *fst.Context) {
		//str, _ := conf.RedisSess.Get("tls:eGhjQnlHbmFTcGlOU0RsemhH")
		//logx.Info(str)
		myData := MyData{}
		if err := ctx.BindPms(&myData); err != nil {
			ctx.FaiErr(err)
			return
		}
		//log.Printf("%v %+v %#v\n", myData, myData, myData)

		myDataT := MyData{}
		_ = ctx.ShouldBindQuery(&myDataT)
		//log.Printf("%v %+v %#v\n", myDataT, myDataT, myDataT)

		ids := ctx.QueryMap("ids")
		firstname := ctx.DefaultQuery("first", "Guest")
		lastname := ctx.Query("last") // shortcut for ctx.Request.URL.Query().Get("lastname")

		message := ctx.PostForm("data")
		nick := ctx.DefaultPostForm("nick", "anonymous")

		//names := ctx.PostFormMap("names")
		ctx.Suc(fst.KV{
			"message": message,
			"nick":    nick,
			"first":   firstname,
			"last":    lastname,
			"ids":     ids,
		})
		//ctx.String(http.StatusOK, fmt.Sprintf("file uploaded!"))
		//ctx.JSON(http.StatusOK, myData)
	})

	// curl -X POST http://127.0.0.1:8099/upload -F "file=@/mnt/d/chende_dev/chende_pri/test.txt"  -H "Content-Type: multipart/form-data"
	bind.Post("/upload", func(ctx *fst.Context) {
		// single file
		file, _ := ctx.FormFile("file")
		log.Println(file.Filename)
		// Upload the file to specific dst.
		if err := ctx.SaveUploadedFile(file, "upload_tst"); err != nil {
			ctx.FaiErr(err)
			return
		}
		ctx.Suc(fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}
