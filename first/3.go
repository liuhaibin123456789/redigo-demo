package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//创建redis连接
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic("redis link is wrong")
		return
	}
	//实例化实现了发布与订阅功能的对象
	subConn := redis.PubSubConn{Conn: conn}
	err = subConn.Subscribe("channel1")
	if err != nil {
		return
	}
	for true {
		//接收频道消息
		switch v := subConn.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
}
