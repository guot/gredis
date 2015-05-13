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

//保存Key值
func (c *Gredis) Set(key string, value string, expireTime int) error {
	cmdstr := buildSetCmd(key, value, expireTime)
	//fmt.Println("向服务器发送命令：", cmdstr)
	_, err := c.conn.Write([]byte(cmdstr))
	str, _, _ := bufio.NewReader(io.Reader(c.conn)).ReadLine()
	fmt.Printf("%s", str)
	return err
}
func (c *Gredis) Get(key string) ([]byte, error) {
	cmdstr := buildGetCmd(key)
	_, err := c.conn.Write([]byte(cmdstr))
	fmt.Printf("send get cmd  is :%s .\n", cmdstr)
	buf := make([]byte, 1024)
	n, _ := c.conn.Read(buf)
	fmt.Printf("get return value is :%s .\n", buf[:n])

	return buf[:n], err
}
func main() {
	client, e := new(Gredis).NewClient("localhost:6379")
	if e != nil {
		log.Fatal(e)
	}
	client.Set("mykey1", "myvalue", -1)
}
