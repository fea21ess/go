package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"hello/controllers"
	"hello/controllers/file"
	"hello/controllers/gdpr"
	"hello/controllers/ldap"
	"hello/controllers/logs"
	"hello/controllers/path"
	"hello/controllers/rce"
	"hello/controllers/redirect"
	"hello/controllers/sql"
	"hello/controllers/ssrf"
	"hello/controllers/xpath"
	"hello/controllers/xss"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/vuln/sql", &sql.IndexSqlController{})
	beego.Router("/vuln/sql/orm", &sql.OrmController{})
	beego.Router("/vuln/sql/db", &sql.SqlDbController{})
	beego.Router("/vuln/sql/mongdb", &sql.MongSqlController{})
	beego.Router("/vuln/rce", &rce.CmdIndexController{})
	beego.Router("/vuln/rce/exec", &rce.CmdExecController{})
	beego.Router("/vuln/rce/pexec", &rce.PipelineExecController{})
	beego.Router("/vuln/file", &file.FileIndexController{})
	beego.Router("/vuln/file/read", &file.ReadController{})
	beego.Router("/vuln/file/remove", &file.RemoveController{})
	beego.Router("/vuln/file/removeAll", &file.RemoveAllController{})
	beego.Router("/vuln/path", &path.PathIndexController{})
	beego.Router("/vuln/path/writeHello", &path.OpenFilePathController{})
	beego.Router("/vuln/xss", &xss.IndexXssController{})
	beego.Router("/vuln/xss/fmt", &xss.FmtXssController{})
	beego.Router("/vuln/xss/beego", &xss.BeegoXssController{})
    beego.Router("/vuln/ldap",&ldap.IndexLdapController{})
	beego.Router("/vuln/ldap/login",&ldap.LoginLdapController{})
	beego.Router("/vuln/ldap/inject",&ldap.InjectLdapController{})
	beego.Router("/vuln/xpath",&xpath.IndexXpathController{})
	beego.Router("/vuln/xpath/login",&xpath.LoginXpathController{})
    beego.Router("/vuln/gdpr",&gdpr.IndexGdprController{})
    beego.Router("/vuln/redirect",&redirect.IndexRedirectController{})
	beego.Router("/vuln/redirect/send",&redirect.UrlRedirectController{})
	beego.Router("/vuln/ssrf",&ssrf.IndexSsrfController{})
	beego.Router("/vuln/ssrf/send",&ssrf.UrlSsrfController{})
	beego.Router("/vuln/logs",&logs.IndexLogsController{})
	beego.Router("/vuln/logs/save",&logs.SaveLogsController{})
}
