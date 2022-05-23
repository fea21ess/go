package sql

import beego "github.com/beego/beego/v2/server/web"

type IndexSqlController struct {
	beego.Controller
}

func (c *IndexSqlController) Get() {

	c.TplName = "vuln/sql/index.tpl"
}