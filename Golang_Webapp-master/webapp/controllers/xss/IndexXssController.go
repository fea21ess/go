package xss

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type IndexXssController struct {
	beego.Controller
}

func (c *IndexXssController) Get() {
	c.TplName = "vuln/xss/index.tpl"
}

type FmtXssController struct {
	beego.Controller
}

func (c *FmtXssController) Get() {
	name := c.GetString("name")
	age := c.GetString("age")
	fmt.Fprintf(c.Ctx.ResponseWriter, "<html><head><title>test</title></head><body>\n")
	fmt.Fprintf(c.Ctx.ResponseWriter, "%s is %s years old.\n", name, age)
	fmt.Fprintf(c.Ctx.ResponseWriter, "</body></html>")
}

type BeegoXssController struct {
	beego.Controller
}

func (c *BeegoXssController) Get() {
	name := c.GetString("name")
	age := c.GetString("age")
	c.Ctx.WriteString(name+" is "+age+" years old.\n")
}