<html>
<head>
    <title>OrmShowView</title>
</head>
<body>

<div>
    {{if .err}}
    <div>
        删除成功
    </div>
    {{.else}}
    <div>
        删除失败
    </div>
    {{end}}
</div>
</body>
</html>