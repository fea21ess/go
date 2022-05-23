package redirect

import beego "github.com/beego/beego/v2/server/web"

type IndexRedirectController struct {
	beego.Controller
}

func (c *IndexRedirectController) Get() {
	c.TplName = "vuln/redirect/index.tpl"
}

type UrlRedirectController struct {
	beego.Controller
}

func (c *UrlRedirectController) Get() {
	url := c.GetString("url")
	c.Redirect(url,302)
}