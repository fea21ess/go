<html>
<head>
    <title>XpathIndexView</title>
</head>
<body>
<h3>Xpath登录框</h3>
<div>
    <form action="{{.url}}/login" method="post">
        <p>username: <input type="text" name="username" /></p>
        <p>password: <input type="password" name="password" /></p>
        <input type="submit" value="Submit" />
    </form>
    {{if .error}}
    <div style="color: red;">
        {{.error}}
    </div>
    {{end}}
</div>
</body>
</html>