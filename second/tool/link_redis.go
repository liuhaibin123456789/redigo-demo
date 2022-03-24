package tool

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var RDS redis.Conn //RDS redis连接

func LinkRedis() error {
	conn, err := redis.Dial("tcp", ":6379", redis.DialWriteTimeout(10*time.Second), redis.DialReadTimeout(10*time.Second))
	if err != nil {
		panic(err)
		return err
	}
	RDS = conn
	return err
}

func Close(conn redis.Conn) error {
	return conn.Close()
}
