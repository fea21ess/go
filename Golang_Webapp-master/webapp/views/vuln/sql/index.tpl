<html>
<head>
    <title>SqlIndexView</title>
</head>
<body>
<h3>Sql注入漏洞靶场</h3>
<ul>
    <li>github.com/beego/beego/v2/client/orm</li>
    <span>o.Raw(sql).Values():<a target="_blank" href="./sql/orm?id=1">./sql/orm?id=1</a></span>
    <li>database/sql</li>
    <span>db.QueryRow(sql):<a target="_blank" href="./sql/db?id=1&vul=QueryRow">./sql/db?id=1&vul=QueryRow</a></span>
    <li>database/sql</li>
    <span>db.Query(sql):<a target="_blank" href="./sql/db?id=1&vul=Query">./sql/db?id=1&vul=Query</a></span>
    <li>database/sql</li>
    <span>db.QueryContext(sql):<a target="_blank" href="./sql/db?id=1&vul=QueryContext">./sql/db?id=1&vul=QueryContext</a></span>
    <li>database/sql</li>
    <span>db.Exec(sql):<a target="_blank" href="./sql/db?id=1&vul=Exec&email=root@admin.com">./sql/db?id=1&vul=Exec&email=root@admin.com</a></span>
    <li>database/sql</li>
    <span>db.ExecContext(sql):<a target="_blank" href="./sql/db?id=1&vul=ExecContext&email=root@admin.com">./sql/db?id=1&vul=ExecContext&email=root@admin.com</a></span>
</ul>

<h3>NoSql注入Mongodb靶场</h3>
<ul>
    <li>go.mongodb.org/mongo-driver/mongo</li>
    <span>collection.Find:<a target="_blank" href="./sql/mongdb?username=admin">./sql/mongdb?username=admin</a></span>
    <li>go.mongodb.org/mongo-driver/mongo</li>
    <span>collection.FindOne:<a target="_blank" href="./sql/mongdb?username=admin&flag=true">./sql/mongdb?username=admin&flag=true</a></span>
</ul>

<h3>Sql注入Postgresql靶场(暂不支持)</h3>
<ul>
    <li>github.com/jackc/pgx/v4</li>
    <span>conn.Query(sql):<a target="_blank" href="./sql/db?id=1&vul=pgx">./sql/db?id=1&vul=pgx</a></span>
</ul>
</body>
</html>