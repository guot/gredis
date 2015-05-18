package gredis

import (
	"strings"
	"testing"
)

var SAddTests = []struct {
	id    int
	key   string
	value string

	exp string
}{
	{1, "mkey", "mvalue", "*3\r\n$4\r\nsadd\r\n$4\r\nmkey\r\n$6\r\nmvalue\r\n"},
}

func TestBuildSAddCmd(t *testing.T) {
	for _, tcast := range SAddTests {
		val := buildSAddCmd(tcast.key, tcast.value)
		if !strings.EqualFold(val, tcast.exp) {
			t.Errorf("测试%d没有通过，期望值为：%s,运行时值为：%s", tcast.id, tcast.exp, val)
		}
	}

}

var SPopTests = []struct {
	id  int
	key string

	exp string
}{
	{1, "mkey", "*3\r\n$4\r\nspop\r\n$4\r\nmkey\r\n$2\r\n10\r\n"},
}

func TestBuildSPopCmd(t *testing.T) {
	for _, tcast := range SPopTests {
		val := buildSPopCmd(tcast.key)
		if !strings.EqualFold(val, tcast.exp) {
			t.Errorf("测试%d没有通过，期望值为：%s,运行时值为：%s", tcast.id, tcast.exp, val)
		}
	}

}
