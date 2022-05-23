package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var books []map[string]interface{}

	articleStrings := `[
							{"name":"001 - SQL注入漏洞","path":"/vuln/sql"},
							{"name":"002 - RCE命令执行","path":"/vuln/rce"},
							{"name":"003 - File文件读取/删除","path":"/vuln/file"},
							{"name":"004 - 路径穿越","path":"/vuln/path"},
							{"name":"005 - XSS跨站脚本","path":"/vuln/xss"},
							{"name":"006 - LDAP注入","path":"/vuln/ldap"},
							{"name":"007 - Xpath注入","path":"/vuln/xpath"},
							{"name":"008 - GDPR敏感信息泄露","path":"/vuln/gdpr"},
							{"name":"009 - Redirect重定向","path":"/vuln/redirect"},
							{"name":"010 - SSRF漏洞","path":"/vuln/ssrf"},
							{"name":"011 - Log注入","path":"/vuln/logs"}
						]`

	books = getJson(articleStrings)
	c.Data["rows1"] = books
	c.TplName = "index.tpl"
}

func getJson(articleStrings string) []map[string]interface{}{
	var articleSlide []map[string]interface{}
	multiErr := json.Unmarshal([]byte(articleStrings),&articleSlide)
	if multiErr!=nil{
		panic("转换出错：")
	}
	return articleSlide
}
