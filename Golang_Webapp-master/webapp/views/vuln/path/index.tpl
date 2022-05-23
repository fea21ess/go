<html>
<head>
    <title>PathIndexView</title>
</head>
<body>
<h3>os.OpenFile目录穿越靶场</h3>
<ul>
    <li>往系统临时目录的filename文件写入"Beego,Hello!"(文件写入、读取主动插桩中统一归到目录穿越)</li>
    <span>os.OpenFile():<a target="_blank" href="./path/writeHello?filename=t.txt">./path/writeHello?filename=t.txt</a></span>
</ul>
</body>
</html>