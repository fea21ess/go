<html>
<head>
    <title>LdapIndexView</title>
</head>
<body>
<h3>LDAP登录框</h3>
<div>
    <span>查询的语句 - filter := "(&(userPassword="+password+")(ou="+username+"))"</span>
    <span>账号密码皆输入*即可绕过登录</span>
    <span>正确的账号密码：administrator/root</span>
    <form action="inject" method="post">
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