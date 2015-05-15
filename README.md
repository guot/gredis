# golang-gredis-client

	基于Redis 2.6.12以上版本开发，

>已实现接口 <br/>
>**SET、GET、DEL**



### API
	Gredis
 
	|Set key value [EX seconds] 
	
		key:待保存的key值，用于唯一标识 待保存 的值 Valeu
		valeu:待保存的值
		EX seconds:设置过期时间 单位是秒 
	返回值
		当保存成功进返回OK
		当NX或XX选项条件不为真时，执行命令返回 Nil
	|Get key
	获取返回值
		key:保存key时所设置的值
	|DEL key,key,key...
	删除指定的Keyw值
		Key:待删除的Key.
		 
## 示例代码
	import "gredis"
	client, e := new(Gredis).NewClient("localhost:6379")
	if e != nil {
		t.Logf("建立连接失败", e.Error())
	}
	err := client.Set("mykey1", "myvalue", -1)
	if err != nil {
		t.Logf("发送命令失败", err.Error())
	}
	res, err := client.Get("mykey1")
	if err != nil {
		t.Logf("发送命令失败", err.Error())
	}
	if bytes.Compare(res, []byte("myvalue")) != 0 {
		t.Logf("返回结果错误,期望值：myvalue．实际值，%s", res)
	}


