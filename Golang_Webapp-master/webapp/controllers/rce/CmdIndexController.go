package rce

import beego "github.com/beego/beego/v2/server/web"

type CmdIndexController struct {
	beego.Controller
}

func (c *CmdIndexController) Get() {
	c.TplName = "vuln/rce/index.tpl"
}