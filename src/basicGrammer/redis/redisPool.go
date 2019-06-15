package redis

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

/*
开源库redigo的使用:
github地址：
https://github.com/garyburd/redigo //go get github.com/garyburd/redigo/redis
文档地址：
http://godoc.org/github.com/garyburd/redigo/redis
*/

var (
	pool *redis.Pool
	redisServer=flag.String("redisServer","127.0.0.1:6381","")
)

func newPool(addr string) *redis.Pool{
	return &redis.Pool{
		MaxIdle:3,
		IdleTimeout:240*time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp",addr)
		},
	}
}

//https://godoc.org/github.com/garyburd/redigo/redis#Script.Do   开源库redigo文档说明
func TestRedis(){
	//https://studygolang.com/pkgdoc  #flag包实现了命令行参数的解析
	flag.Parse()

	pool=newPool(*redisServer)

	conn:=pool.Get()

	defer conn.Close()

	key1:="netname"

	key2:="invitenet"

	value:="{name:syf,age:30}"

	exists,err:=redis.Bool(conn.Do("EXISTS",key1,key2)) //exists结果true/false

	if err!=nil{
		fmt.Println("handle error return from c.Do or type conversion error")
	}

	fmt.Println("exists##",exists)

	//key值不存在，则设置key值
	if exists==false{
		//设置值
		rs,err:=conn.Do("HSET",key1,key2,value)
		if err!=nil{
			fmt.Println("hset错误##",err)
		}
		fmt.Println("rs设置结果##",rs)
	}

	//取值
	rs,err:=redis.Bytes(conn.Do("HGET",key1,key2))
	if err!=nil{
		fmt.Println("hget错误##",err)
	}
	fmt.Println("rs取结果值##",string(rs)) //rs取结果值## {name:syf,age:30}
}

func TestRedisPubAndSub(){

	flag.Parse()

	pool=newPool(*redisServer)

	conn:=pool.Get()

	defer conn.Close()

	channel:="syf"  //订阅/发布的通道

	message:="Publish and Subscribe" //订阅/发布的消息

	conn.Send("SUBSCRIBE",channel)

	conn.Flush()

	reply,err:=conn.Receive()
	if err!=nil{
		fmt.Println("conn.Receive###",err)
	}

	fmt.Println("reply###",reply)

	rs,error:=conn.Do("PUBLISH",channel,message)

	if error!=nil{
		fmt.Println("PUBLISH error",error)
	}

	fmt.Println("PUBLISH rs##",rs)

}

//订阅
func RedisSubscribe(){
	flag.Parse()

	pool=newPool(*redisServer)

	conn:=pool.Get()

	defer conn.Close()

	channel:="syf"

	psc:=redis.PubSubConn{Conn:conn}

	psc.Subscribe(channel) //订阅通道

    for{
		switch v:=psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s:message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println("v##",v)
			return
		}
	}
}

//发布
func RedisPublish(){

	flag.Parse()

	pool=newPool(*redisServer)

	conn:=pool.Get()

	defer conn.Close()

	channel:="syf"

	message:="kaixinyufeng"

	_,err:=conn.Do("PUBLISH",channel,message) //向通道发布数据

	if err!=nil{
		fmt.Println("publish error##",err)
	}
}

/*
参考博客:https://www.cnblogs.com/wdliu/p/9330278.html
*/
func TestRedisSubAndPub(){
	go RedisSubscribe()
	go RedisPublish()
	time.Sleep(time.Second*3)
	/*
	执行结果:
	syf: subscribe 1
	syf:message: kaixinyufeng
	*/
}