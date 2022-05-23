package gdpr

import beego "github.com/beego/beego/v2/server/web"

type IndexGdprController struct {
	beego.Controller
}

func (c *IndexGdprController) Get() {
	c.TplName = "vuln/gdpr/index.tpl"
}
