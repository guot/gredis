// base
package gredis

import (
	"bytes"
	"strconv"
)

var (
	cmdSuffix []byte = []byte("\r\n")
	//命令的前缀
	cmdPrefix byte = '$'
	//命令的后缀
	startPrefix byte = '*'
)

//命令基类
type Cmd struct {
	name    string
	buf     bytes.Buffer
	cmds    []string
	pamCnt  int
	index   int
	execute interface{}
}

//添加命令参数
//如
//c:=newStringCmd("set")
//c.AddCmd("mkey")
//c.AddCmd("mvalue")
func (c *Cmd) addCmd(parm string) {

	c.cmds = append(c.cmds, parm)

}

//获取字符串长度的字符类型
func lenStr(val string) string {
	return strconv.Itoa(len(val))
}

//生成命令文本
func (c *Cmd) toString() string {
	//插入命令长度
	c.buf.WriteByte('*')
	c.buf.Write([]byte(strconv.Itoa(len(c.cmds) + 1)))
	c.buf.Write(cmdSuffix)
	c.buf.WriteByte(cmdPrefix)
	c.buf.Write([]byte(strconv.Itoa(len(c.name))))
	c.buf.Write(cmdSuffix)
	c.buf.Write([]byte(c.name))
	c.buf.Write(cmdSuffix)
	for _, cmdp := range c.cmds {
		//添加　$符号
		c.buf.WriteByte(cmdPrefix)
		//添加　命令长度
		c.buf.Write([]byte(strconv.FormatInt(int64(len(cmdp)), 10)))
		//添加　换行
		c.buf.Write(cmdSuffix)
		//添加　$符号
		//添加　命令
		c.buf.Write([]byte(cmdp))
		//添加　换行
		c.buf.Write([]byte(cmdSuffix))
	}
	return c.buf.String()

}
