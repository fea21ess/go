package path

import (
	beego "github.com/beego/beego/v2/server/web"
	"os"
)

type PathIndexController struct {
	beego.Controller
}

func (c *PathIndexController) Get() {
	c.TplName = "vuln/path/index.tpl"
}

type OpenFilePathController struct {
	beego.Controller
}

func (c *OpenFilePathController) Get() {
	temp := os.TempDir()
	filename := temp + "\\" + c.GetString("filename")
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		c.Ctx.WriteString("发生错误")
	}
	_, _ = f.WriteString("Beego,Hello!")
	c.Ctx.WriteString("写入成功")
	_ = f.Close()
}
