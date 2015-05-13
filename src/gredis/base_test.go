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

func BenchmarkCommand(b *testing.B) {
	var c *Cmd
	for i := 0; i < b.N; i++ {
		c = &Cmd{}
		c.addCmd("mykey")
		c.addCmd("myvalue")
		c = nil
	}
}
