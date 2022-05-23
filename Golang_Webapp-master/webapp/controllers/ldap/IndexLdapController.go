package ldap

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/go-ldap/ldap/v3"
)

type IndexLdapController struct {
	beego.Controller
}

func (c *IndexLdapController) Get() {
	c.TplName = "vuln/ldap/index.tpl"
}

type LoginLdapController struct {
	beego.Controller
}

func (c *LoginLdapController) Get() {
	c.TplName = "vuln/ldap/login.tpl"
}

type InjectLdapController struct {
	beego.Controller
}

func (c *InjectLdapController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	rs := ExampleSearch(username,password)
	if rs {
		c.Data["username"] = username
		c.TplName = "vuln/ldap/inject.tpl"
	}else {
		c.Data["error"] = "登录账号或密码错误"
		c.TplName = "vuln/ldap/login.tpl"
	}

}


func ExampleSearch(username string,password string) bool{
	l, err := ldap.DialURL("ldap://172.18.0.5:389")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()

	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: "cn=admin,dc=moresec,dc=cn",
		Password: "123456",
	})
	if err != nil {
		fmt.Printf("Failed to bind: %s\n\n", err)
	}

	filter := "(&(userPassword="+password+")(ou="+username+"))"
	fmt.Print("exec filter:"+filter)
	searchRequest := ldap.NewSearchRequest(
		"dc=moresec,dc=cn", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, // The filter to apply
		[]string{"dn","cn","uid","ou","userPassword"},                    // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err)
	}
	if len(sr.Entries) != 0 {
		uname := sr.Entries[0].GetAttributeValue("ou")
		pwd := sr.Entries[0].GetAttributeValue("userPassword")
		if uname != "" && pwd != ""{
			return true
		}
	}
	return false
}