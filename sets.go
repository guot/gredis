// gredis
package gredis

//加入缓存　key :缓存的键，value:缓存的值，
//	　　expireTime:缓存的过期时间 单位秒，－1则不过期
//构造Set命令
func buildSAddCmd(key string, value string) string {
	cmd := &Cmd{name: "sadd"}
	cmd.addCmd(key)
	cmd.addCmd(value)
	return cmd.toString()
}

//构造SPOP命令 从集合中弹出前１０个元素
func buildSPopCmd(key string) string {
	cmd := &Cmd{name: "spop"}
	cmd.addCmd(key)
	return cmd.toString()
}
