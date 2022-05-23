<html>
<head>
    <title>OrmShowView</title>
</head>
<body>

{{if .userdata}}
    <div>
        <tr>
            <td>Username:{{.userdata.UserName}}</td>
            <br/>
            <td>Pasword:{{.userdata.PassWord}}</td>
        </tr>
    </div>
{{else}}
    {{range $k,$v := .results}}
    <div>
        <tr>
            <td>Username:{{$v.UserName}}</td>
            <br/>
            <td>Pasword:{{$v.PassWord}}</td>
        </tr>
    </div>
    {{end}}
{{end}}
</body>
</html>