<html>
<head>
    <title>XssIndexView</title>
</head>
<body>
<h3>XSS跨站脚本靶场</h3>
<ul>
    <li>fmt.Fprint输出内容</li>
    <span>fmt.Fprint():<a target="_blank" href="./xss/fmt?name=chen&age=18">./xss/fmt?name=chen&age=18</a></span>
    <li>Beego框架的Ctx.WriteString输出</li>
    <span>fmt.Fprint():<a target="_blank" href="./xss/beego?name=chen&age=18">./xss/beego?name=chen&age=18</a></span>

</ul>
</body>
</html>