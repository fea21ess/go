package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/jackc/pgx/v4"
	beego "github.com/beego/beego/v2/server/web"
	"os"
)

type SqlDbController struct {
	beego.Controller
}

func checkErr(err error) {
	if err != nil {
		log := logs.NewLogger()
		logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
		log.EnableFuncCallDepth(true)
		log.Error(err.Error())
	}
}

func (c *SqlDbController) Get() {
	id := c.GetString("id")
	vul := c.GetString("vul")
	user := &User{}
	db, err := sql.Open("mysql", "root:root@tcp(172.18.0.3:3306)/beego?charset=utf8&loc=Local")
	checkErr(err)
	defer db.Close()
	switch vul {
	case "QueryRow":
		sql := "SELECT id,name,password,email FROM user where id='"+id+"'"
		row := db.QueryRow(sql)
		err = row.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		checkErr(err)
		c.Data["name"] = user.Name
		c.Data["id"] = user.Id
		c.Data["email"] = user.Email
		break
	case "Query":
		sql := "SELECT id,name,password,email FROM user where id='"+id+"'"
		rows, err := db.Query(sql)
		checkErr(err)
		for rows.Next() { //作为循环条件来迭代获取结果集Rows

			//从结果集中获取一行结果
			err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
			checkErr(err)
		}
		defer rows.Close() //关闭结果集，释放链接
		c.Data["name"] = user.Name
		c.Data["id"] = user.Id
		c.Data["email"] = user.Email
		break
	case "QueryContext":
		sql := "SELECT id,name,password,email FROM user where id='"+id+"'"
		ctx, _ := context.WithCancel(context.Background())
		rows, err := db.QueryContext(ctx,sql)
		checkErr(err)
		for rows.Next() { //作为循环条件来迭代获取结果集Rows
			//从结果集中获取一行结果
			err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
			checkErr(err)
		}
		defer rows.Close() //关闭结果集，释放链接
		c.Data["name"] = user.Name
		c.Data["id"] = user.Id
		c.Data["email"] = user.Email
		break
	case "QueryRowContext":
		sql := "SELECT id,name,password,email FROM user where id='"+id+"'"
		ctx, _ := context.WithCancel(context.Background())
		err := db.QueryRowContext(ctx,sql).Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		checkErr(err)
		c.Data["name"] = user.Name
		c.Data["id"] = user.Id
		c.Data["email"] = user.Email
		break
	case "Exec":
		email := c.GetString("email")
		sql := "update user set email='"+email+"' where id='"+id+"'"
		rows,err := db.Exec(sql)
		checkErr(err)
		num, err := rows.RowsAffected()
		if num != 0 {
			c.Data["num"] = num
			c.Data["flag"] = true
		}else {
			c.Data["num"] = -1
			c.Data["flag"] = false
		}
		break
	case "ExecContext":
		email := c.GetString("email")
		sql := "update user set email='"+email+"' where id='"+id+"'"
		ctx, _ := context.WithCancel(context.Background())
		rows, err := db.ExecContext(ctx,sql)
		checkErr(err)
		num, err := rows.RowsAffected()
		if num != 0 {
			c.Data["num"] = num
			c.Data["flag"] = true
		}else {
			c.Data["num"] = -1
			c.Data["flag"] = false
		}

		break
	case "pgx":
		urlExample := "postgres://postgres:root@192.168.44.128:5432/beego"
		conn, err := pgx.Connect(context.Background(), urlExample)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		}
		defer conn.Close(context.Background())

		err = conn.QueryRow(context.Background(), "SELECT id,name,password,email FROM users where id='"+id+"'").Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		}

		c.Data["name"] = user.Name
		c.Data["id"] = user.Id
		c.Data["email"] = user.Email
		break
	default:

	}


	c.TplName = "vuln/sql/OrmShow.tpl"
}