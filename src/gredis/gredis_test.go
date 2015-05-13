// gredis_test
package gredis

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMain(t *testing.T) {
	client, e := new(Gredis).NewClient("localhost:6379")
	if e != nil {
		t.Logf("建立连接失败", e.Error())
	}
	err := client.Set("mykey1", "myvalue", -1)
	if err != nil {
		t.Logf("发送命令失败", err.Error())
	}
}
func TestClientGet(t *testing.T) {
	client, e := new(Gredis).NewClient("localhost:6379")
	if e != nil {
		t.Logf("建立连接失败", e.Error())
	}
	client.Set("mykey1", "myvalue", -1)
	res, err := client.Get("mykey1")
	if err != nil {
		t.Logf("发送命令失败", err.Error())
	}
	if bytes.Compare(res, []byte("mvalue1")) != 0 {
		t.Logf("返回结果错误,期望值：myvalue1．实际值，%s", res)
	}
}
func BenchmarkClient(b *testing.B) {
	b.StopTimer()
	client, e := new(Gredis).NewClient("localhost:6379")
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		if e != nil {
			b.Logf("建立连接失败", e.Error())
		}
		err := client.Set(strconv.Itoa(i), "myvalue", -1)
		if err != nil {
			b.Logf("发送命令失败", err.Error())
		}
	}
}
