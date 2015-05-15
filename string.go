// string
package gredis

import (
	"strconv"
)

//加入缓存　key :缓存的键，value:缓存的值，
//	　　expireTime:缓存的过期时间 单位秒，－1则不过期
//构造Set命令
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

//构造Get 命令
func buildGetCmd(key string) string {
	cmd := &Cmd{name: "get"}
	cmd.addCmd(key)
	return cmd.toString()
}

//构造Del 命令
func buildDelCmd(keys ...string) string {
	cmd := &Cmd{name: "del"}
	for _, key := range keys {
		cmd.addCmd(string(key))
	}
	return cmd.toString()
}
