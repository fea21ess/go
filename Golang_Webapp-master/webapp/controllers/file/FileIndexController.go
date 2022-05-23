package file

import beego "github.com/beego/beego/v2/server/web"

type FileIndexController struct {
	beego.Controller
}

func (c *FileIndexController) Get() {
	c.TplName = "vuln/file/index.tpl"
}