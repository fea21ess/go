package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "hello/controllers/filter"
	_ "hello/routers"
	"html/template"
	"net/http"
)

func pageNotFound(rw http.ResponseWriter, r *http.Request){
	t,_:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/404.html")
	data :=make(map[string]interface{})
	data["content"] = "Page Not Found"
	t.Execute(rw, data)
}

func main() {
	beego.InsertFilterChain("/*", func(next beego.FilterFunc) beego.FilterFunc {
		return func(ctx *context.Context) {
			// do something
			log := logs.NewLogger()
			logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
			log.EnableFuncCallDepth(true)
			log.Debug("Logs:Hello")
			// don't forget this

			next(ctx)

			// do something
		}
	})

	beego.ErrorHandler("404", pageNotFound)
	beego.Run()
}

