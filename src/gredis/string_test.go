// string
package gredis

import (
	"strings"
	"testing"
)

var setTests = []struct {
	id      int
	key     string
	value   string
	expTime int
	exp     string
}{
	{1, "mkey", "mvalue", -1, "*3\r\n$3\r\nset\r\n$4\r\nmkey\r\n$6\r\nmvalue\r\n"},
	{2, "mkey", "mvalue", 1000, "*5\r\n$3\r\nset\r\n$4\r\nmkey\r\n$6\r\nmvalue\r\n$2\r\nEX\r\n$4\r\n1000\r\n"},
}

func TestBuildSetCmd(t *testing.T) {
	for _, tcast := range setTests {
		val := buildSetCmd(tcast.key, tcast.value, tcast.expTime)
		if !strings.EqualFold(val, tcast.exp) {
			t.Errorf("测试%d没有通过，期望值为：%s,运行时值为：%s", tcast.id, tcast.exp, val)
		}
	}

}

func TestBuildGetCmd(t *testing.T) {

}
