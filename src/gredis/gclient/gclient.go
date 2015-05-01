// gclient
package main

import (
	"bytes"
	"fmt"
	"net"
	"utils/strings"
) 
type Gredis struct {
	host string
	conn net.Conn
}
type Cmmand struct {
	cmd string
}
const  	get,set,put,add,del="get","set","put","add","del"	
 
/**
enter here
*/
func (c *Gredis) NewClient(host string) (client *Gredis, e error) {
	client = &Gredis{host: host}
	e = c.connect()
	return
}
func (c *Gredis) connect() error {
	var err error
	c.conn, err = net.Dial("tcp", c.host)
	return err
}
func (c *Gredis) Set(key, value string) error {
	var bf bytes.Buffer

	bf.Write([]byte("*3\r\nasdfsdf"))
	bf.WriteString("*3\r\n")
	bf.WriteString("$3\r\n")
	bf.WriteString("$set\r\n")

	bf.WriteString("$" + sutils.ValueOf(len(value)) + "\r\n")
	bf.WriteString("$" + value + "\r\n")
	fmt.Print(bf.String())
	return nil
}
func buildCommand(stype string , data map[string]string){
	switch stype   {
		case  get:
			fmt.Printf("the command  is get")		
	}
}
func (c *Gredis) Get(key string) (string, error) {
	var bf bytes.Buffer
	sutils.AppendToNewLine(&bf,"*2")
	bf.WriteString("*2\r\n")
	bf.WriteString("$3\r\n")
	bf.WriteString("$get\r\n")

	bf.WriteString("$" + sutils.ValueOf(len(key)) + "\r\n")
	bf.WriteString("$" + key + "\r\n")
	fmt.Print(bf.String())
	return bf.String(), nil
}
func main() {
	client, _ := new(Gredis).NewClient("localhost:8311")
	client.Set("mkey", "myvalue")
	buildCommand(get,nil)
//	val, _ := client.Get("mkey")
//	fmt.Print(val)
}
