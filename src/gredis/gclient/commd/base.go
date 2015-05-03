// base
package commd

import (
	"bufio"
	"strconv"
)

var (
	cmdsuffix   []byte = []byte("\r\n")
	cmdPrefix   byte   = '$'
	startPrefix byte   = '*' 
	SetCmd = newStringCmd("set")
)

type cmd struct {
	bufio.Writer
	cmds    [10]string
	pamCnt  int
	cmdType string
	index int
}

func newStringCmd(name string) *cmd {
	cmd := &cmd{cmdType: "String"}
	cmd.WriteByte(startPrefix)
 
	cmd.cmd[cmd.index++]=name
	return cmd
}
func (c *cmd) AddCmd(parm string) {
	c.WriteByte(cmdPrefix)
	c.Write([]byte(parm))

}
func (c *cmd) toString() string {
	//插入命令长度
	c.WriteByte(strconv.Itoa(len(c.cmds)))
	for cmd := range c.cmds[1:] {
		cmd.Write(cmd)
		cmd.Write([]byte(cmdsuffix))
	}
	return cmd.toString()
}
