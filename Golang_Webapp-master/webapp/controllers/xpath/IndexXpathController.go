package xpath

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gopkg.in/xmlpath.v2"
	"strings"
)

type IndexXpathController struct {
	beego.Controller
}

func (c *IndexXpathController) Get() {
	c.Data["url"] = c.Ctx.Request.URL
	c.TplName = "vuln/xpath/login.tpl"
}

type LoginXpathController struct {
	beego.Controller
}

func (c *LoginXpathController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	flag := CheckLoginUser(username,password)
	if flag {
		c.Data["username"] = username
		c.TplName = "vuln/xpath/main.tpl"
	}else {
		c.Data["error"] = "登录失败,请检查账号密码"
		c.Data["url"] = strings.Replace(c.Ctx.Request.URL.String(),"/login","",1)

		c.TplName = "vuln/xpath/login.tpl"
	}

}

func CheckLoginUser(username string, password string) bool{
	xmlDoc := `<?xml version="1.0" encoding="UTF-8"?>
 
				<bookstore>
				 
					<user role="1" username="admin" password="admin888"></user>
					<user role="1" username="root" password="root123"></user>
					<user role="2" username="guest" password="123456"></user>
				  
				</bookstore>`
	reader := strings.NewReader(xmlDoc)
	rs, err := xmlpath.ParseHTML(reader)
	if err != nil {
		fmt.Println("Error: "+err.Error())
		return false
	} else {
		path := xmlpath.MustCompile("//bookstore/user[@username='"+username+"' and @password='"+password+"']")
		_,flag := path.String(rs)
		fmt.Println(flag)
		return flag
	}
}