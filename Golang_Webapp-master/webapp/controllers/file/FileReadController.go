package file

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"os"
)

type ReadController struct {
	beego.Controller
}

func (c *ReadController) Get() {
	filename := c.GetString("filename")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data["rs"] = string(buffer)
	c.TplName = "vuln/file/read.tpl"
}

type RemoveController struct {
	beego.Controller
}

func (c *RemoveController) Get() {
	filename := c.GetString("filename")
	err := os.Remove(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data["err"] = true
	c.TplName = "vuln/file/remove.tpl"
}

type RemoveAllController struct {
	beego.Controller
}

func (c *RemoveAllController) Get() {
	filename := c.GetString("filename")
	err := os.RemoveAll(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data["err"] = true
	c.TplName = "vuln/file/remove.tpl"
}