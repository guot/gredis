# golang-gredis-client

	基于Redis 2.6.12以上版本开发，

>准备实现接口 <br/>
>**SET、GET**

### 使用说明
>1、配置conf目录下的conf.properties文件。<br>
>  _#以实际情况进行配置_

redis.server.host=192.168.1.1

### Example
	var client=Gclient
 
	SET key value [EX seconds] [PX milliseconds] [NX|XX]
	
		key:待保存的key值，用于唯一标识 待保存 的值 Valeu
		valeu:待保存的值
		EX seconds:设置过期时间 单位是秒
		PX milliseconds: 设置过期时间 单位是毫秒
		NX:只有key不存在时才保存
		XX:只有key存在时才保存
	返回值
		当保存成功进返回OK
		当NX或XX选项条件不为真时，执行命令返回 Nil
	 GET key
	获取返回值
		key:保存key时所设置的值


