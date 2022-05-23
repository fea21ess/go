package sql

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type OrmController struct {
	beego.Controller
}

func init(){
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(172.18.0.3:3306)/beego?charset=utf8&loc=Local")

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}

//SQL Injection: http://localhost:8080/sql/orm?id=1%27%20and%201=1%20%23
func (c *OrmController) Get() {
	id := c.GetString("id")
	o := orm.NewOrm()

	var maps []orm.Params
	//var user User
	num, err := o.Raw("SELECT * FROM user where id='"+id+"'").Values(&maps)
	//num, err := o.Raw("SELECT * FROM user where id='"+id+"'",1).QueryRows(&user)
	if err == nil && num > 0 {
		c.Data["id"] = maps[0]["id"]
		c.Data["name"] = maps[0]["name"]
		c.Data["email"] = maps[0]["email"]
	}

	c.TplName = "vuln/sql/OrmShow.tpl"
}
