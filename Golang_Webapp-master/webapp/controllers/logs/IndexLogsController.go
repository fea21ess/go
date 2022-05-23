package logs

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type IndexLogsController struct {
	beego.Controller
}

func (c *IndexLogsController) Get() {
	c.TplName = "vuln/logs/index.tpl"
}

type SaveLogsController struct {
	beego.Controller
}

// beego 日志配置结构体
type LoggerConfig struct {
	FileName            string `json:"filename"`
	Level               int    `json:"level"`    // 日志保存的时候的级别，默认是 Trace 级别
	Maxlines            int    `json:"maxlines"` // 每个文件保存的最大行数，默认值 1000000
	Maxsize             int    `json:"maxsize"`  // 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	Daily               bool   `json:"daily"`    // 是否按照每天 logrotate，默认是 true
	Maxdays             int    `json:"maxdays"`  // 文件最多保存多少天，默认保存 7 天
	Rotate              bool   `json:"rotate"`   // 是否开启 logrotate，默认是 true
	Perm                string `json:"perm"`     // 日志文件权限
	RotatePerm          string `json:"rotateperm"`
	EnableFuncCallDepth bool   `json:"-"` // 输出文件名和行号
	LogFuncCallDepth    int    `json:"-"` // 函数调用层级
}

func (c *SaveLogsController) Get() {
	var logCfg = LoggerConfig{
		FileName:            "./daily.log",
		Level:               7,
		EnableFuncCallDepth: true,
		LogFuncCallDepth:    3,
		RotatePerm:          "777",
		Perm:                "777",
	}

	logName := c.GetString("logname")
	// 设置beego log库的配置
	b, _ := json.Marshal(&logCfg)
	logs.SetLogger(logs.AdapterFile, string(b))
	logs.EnableFuncCallDepth(logCfg.EnableFuncCallDepth)
	logs.SetLogFuncCallDepth(logCfg.LogFuncCallDepth)
	logs.Info("SaveLogsController Process start in "+logName)
	c.TplName = "vuln/logs/save.tpl"
}