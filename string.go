// string
package gredis

import (
	"strconv"
)

//加入缓存　key :缓存的键，value:缓存的值，
//	　　expireTime:缓存的过期时间 单位秒，－1则不过期

func buildSetCmd(key string, value string, expireTime int) string {
	cmd := &Cmd{name: "set"}
	cmd.addCmd(key)
	cmd.addCmd(value)
	if expireTime != -1 {
		cmd.addCmd("EX")
		cmd.addCmd(strconv.Itoa(expireTime))
	}
	return cmd.toString()
}
func buildGetCmd(key string) string {
	cmd := &Cmd{name: "get"}
	cmd.addCmd(key)
	return cmd.toString()
}
