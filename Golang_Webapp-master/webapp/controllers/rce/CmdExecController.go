package rce

import (
	"bufio"
	"bytes"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io"
	"os"
	"os/exec"
	"strings"
)

type CmdExecController struct {
	beego.Controller
}

type PipelineExecController struct {
	beego.Controller
}

func (c *PipelineExecController) Get() {
	cmd := c.GetString("cmd")
	rs := CmdAndChangeDirToShow(cmd)
	fmt.Print(rs)
	c.Data["rs"] = rs
	c.TplName = "vuln/rce/result.tpl"
}

func (c *CmdExecController) Get() {
	cmd := c.GetString("cmd")
	rs, _ := Cmd(cmd)
	fmt.Print(rs)
	c.Data["rs"] = rs
	c.TplName = "vuln/rce/result.tpl"
}

func Cmd(commandName string) (string, error) {
	cmd := exec.Command(commandName)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

func CmdAndChangeDirToShow(commandName string) string {
	cmd := exec.Command(commandName)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe: ", err)
		return ""
	}
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		fmt.Println("cmd.Start: ", err)
		return ""
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	var b strings.Builder
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		b.WriteString(line)
	}
	err = cmd.Wait()

	return b.String()
}