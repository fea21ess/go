<html>
<head>
    <title>OrmShowView</title>
</head>
<body>
{{if .num}}
<div>
    {{if .flag}}
        <p>Your data is change.</p>
    {{else}}
        <p>Your data not change.</p>
    {{end}}
</div>
{{else}}
<div>
    <p>Your ID:{{.id}}</p>
    <br/>
    <p>Your Name:{{.name}}</p>
    <br/>
    <p>Your Email:{{.email}}</p>
</div>
{{end}}
</body>
</html>