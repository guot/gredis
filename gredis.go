// gclient
package gredis

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Gredis struct {
	strcmd Cmd
	host   string
	conn   net.Conn
}

/**
enter here
*/
func (c *Gredis) NewClient(host string) (client *Gredis, e error) {
	c = &Gredis{host: host}
	e = c.connect()
	return c, e
}
func (c *Gredis) connect() error {
	var err error
	c.conn, err = net.Dial("tcp", c.host)
	return err
}
func (c *Gredis) Close() {
	c.conn.Close()
}

//保存Key值
func (c *Gredis) Set(key string, value string, expireTime int) error {
	cmdstr := buildSetCmd(key, value, expireTime)
	//fmt.Println("向服务器发送命令：", cmdstr)
	_, err := c.conn.Write([]byte(cmdstr))
	str, _, _ := bufio.NewReader(io.Reader(c.conn)).ReadLine()
	log.Printf("Set return val:%s\r\n", str)
	return err
}
func (c *Gredis) Get(key string) ([]byte, error) {
	cmdstr := buildGetCmd(key)
	_, err := c.conn.Write([]byte(cmdstr))
	log.Printf("send get cmd  is :%s .\n", cmdstr)
	buf := make([]byte, 1024)
	n, _ := c.conn.Read(buf)
	val := pareseResp(buf[:n])
	fmt.Printf("get return value is :%s .\n", val)

	return val, err
}
func (c *Gredis) Del(keys ...string) error {
	cmdstr := buildDelCmd(keys...)
	_, err := c.conn.Write([]byte(cmdstr))
	log.Printf("send del cmd  is :%s .\n", cmdstr)
	buf := make([]byte, 1024)
	n, _ := c.conn.Read(buf)
	val := pareseResp(buf[:n])
	//	byteBuf := bytes.NewBuffer(buf[:n])
	//	val, _ := binary.ReadVarint(byteBuf)
	fmt.Printf("get return del value is:%s.\n", val)
	return err
}
func (c *Gredis) SAdd(key string, value string) error {
	cmdstr := buildSAddCmd(key, value)
	_, err := c.conn.Write([]byte(cmdstr))
	log.Printf("send sadd cmd is :%s.\n", cmdstr)
	buf := make([]byte, 1024)
	n, _ := c.conn.Read(buf)
	pareseResp(buf[:n])
	return err
}

func (c *Gredis) SPop(key string) ([]byte, error) {
	cmdstr := buildSPopCmd(key)
	_, err := c.conn.Write([]byte(cmdstr))
	log.Printf("send sadd cmd is :%s.\n", cmdstr)
	buf := make([]byte, 1024)
	n, _ := c.conn.Read(buf)
	val := pareseResp(buf[:n])
	return val, err
}
