// base_test
package gredis

import "testing"

var cmdTests = []struct {
	id  int
	cmd []string
	exp string
}{
	{1, []string{"set"}, "*1\r\n$3\r\nset\r\n"},
	{2, []string{"set", "mykey"}, "*2\r\n$3\r\nset\r\n$5\r\nmykey\r\n"},
	{3, []string{"set", "mykey", "myvalue"}, "*3\r\n$3\r\nset\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"},
	{4, []string{"get", "mykey"}, "*2\r\n$3\r\nget\r\n$5\r\nmykey\r\n"},
}

func TestCommand(t *testing.T) {

	var c *Cmd

	for _, ctx := range cmdTests {
		for _, cmd := range ctx.cmd {
			if c == nil {
				c = &Cmd{name: cmd}
			} else {
				c.addCmd(cmd)
			}
		}
		res := c.toString()
		//		t.Logf("Logtext:\r\n" + res)
		if res != ctx.exp {
			t.Errorf("测试%d没有通过，期望值为：%s,运行时值为：%s", ctx.id, ctx.exp, res)
		}
		c = nil
	}
}

var responseTests = []struct {
	id   int
	resp []byte
	exp  string
}{
	{1, []byte("$5\r\nvalue"), "value"},
	{2, []byte("+ok\r\n"), "ok"},
	{3, []byte("-error"), "error"},
	{4, []byte("*0\r\n"), ""},
	{5, []byte(":10\r\n"), "10"},
	{6, []byte("m10\r\n"), "error"},
}

func TestPareseResp(t *testing.T) {

	for _, tcast := range responseTests {
		val := pareseResp(tcast.resp)
		if string(val) != tcast.exp {
			t.Errorf("测试%d没有通过，期望值为：%s,运行时值为：%s", tcast.id, tcast.exp, tcast.resp)
		}
	}

}
func BenchmarkCommand(b *testing.B) {
	var c *Cmd
	for i := 0; i < b.N; i++ {
		c = &Cmd{}
		c.addCmd("mykey")
		c.addCmd("myvalue")
		c = nil
	}
}
