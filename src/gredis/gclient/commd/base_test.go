// base_test
package commd

import "testing"

var cmdTests = []struct {
	id  int
	cmd []string
	exp string
}{
	{1, []string{"set"}, "*1\r\n$3\r\n$set\r\n"},
	{2, []string{"set", "mykey"}, "*2\r\n$3\r\n$set\r\n$5\r\n$mykey\r\n"},
	{3, []string{"set", "mykey", "myvalue"}, "*3\r\n$3\r\n$set\r\n$5\r\n$mykey\r\n$7\r\n$myvalue\r\n"},
}

func TestCommand(t *testing.T) {
	var c *Cmd
	for _, ctx := range cmdTests {
		for _, cmd := range ctx.cmd {
			if c == nil {
				c = newStringCmd(cmd)
			}else{
				c.AddCmd(cmd)
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
