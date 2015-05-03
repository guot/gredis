// base
package commd

import (
	"bytes"
	"strconv"
)

var (
	cmdsuffix []byte = []byte("\r\n")
	//命令的前缀
	cmdPrefix byte = '$'
	//命令的后缀
	startPrefix byte = '*'
	SetCmd           = newStringCmd("set")
)

//命令基类
type cmd struct {
	bytes.Buffer
	cmds    []string
	pamCnt  int
	cmdType string
	index   int
}

//构造新的命令， 传入命令名
func newStringCmd(name string) *cmd {
	cmd := &cmd{cmdType: "String"}
	cmd.WriteByte(startPrefix)

	cmd.cmds = append(cmd.cmds, name)

	return cmd
}

//添加命令参数
//如
//c:=newStringCmd("set")
//c.AddCmd("mkey")
//c.AddCmd("mvalue")
func (c *cmd) AddCmd(parm string) {
	c.WriteByte(cmdPrefix)
	c.Write([]byte(lenStr(parm)))
	c.WriteByte(cmdPrefix)
	c.Write([]byte(parm))
	c.cmds = append(c.cmds, parm)
}

//获取字符串长度的字符类型
func lenStr(val string) string {
	return strconv.FormatInt(int64(len(val)), 10)
}
func (c *cmd) toString() string {
	//插入命令长度
	
	c.Write([]byte(strconv.Itoa(len(c.cmds))))
	for _,cmdp := range c.cmds{
		c.Write([]byte(cmdp))
		c.Write([]byte(cmdsuffix))
	}
	return (c.(*bytes.Buffer)).toString()

}
