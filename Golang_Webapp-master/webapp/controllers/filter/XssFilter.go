package filter

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, filterUser2)
	beego.InsertFilter("/*", beego.BeforeRouter, filterUser)
}

var filterUser = func(ctx *context.Context) {
	log := logs.NewLogger()
	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	log.EnableFuncCallDepth(true)
	log.Debug("Logs:Test")
}

var filterUser2 = func(ctx *context.Context) {
	log := logs.NewLogger()
	logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	log.EnableFuncCallDepth(true)
	log.Debug("Logs:Test2")
}