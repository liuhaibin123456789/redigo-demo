package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis" //该库适合对redis命令熟悉的对象使用
)

func main() {
	conn, err := redis.Dial("tcp", ":6379", redis.DialOption{})
	if err != nil {
		return
	}
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	do, err := conn.Do("set", "k1", "k1")
	if err != nil {
		return
	}
	fmt.Println(do)
	do, err = conn.Do("get", "k1")
	if err != nil {
		return
	}
	fmt.Println(do)
	do, err = redis.String(do, err)
	if err != nil {
		return
	}
	fmt.Println(do)
}
