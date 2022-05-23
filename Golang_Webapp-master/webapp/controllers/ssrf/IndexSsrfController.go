package ssrf

import (
	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
)

type IndexSsrfController struct {
	beego.Controller
}

func (c *IndexSsrfController) Get() {
	c.TplName = "vuln/ssrf/index.tpl"
}

type UrlSsrfController struct {
	beego.Controller
}

func (c *UrlSsrfController) Get() {
	url := c.GetString("url")
	request := httplib.Get(url)
	data, _ := request.String()
	c.Data["data"] = data
	c.TplName = "vuln/ssrf/result.tpl"
}