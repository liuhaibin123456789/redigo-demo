package main

import "github.com/go-redis/redis" //封装了redis库的命令

func main() {
	client := redis.NewClient(&redis.Options{DB: 0, Addr: "localhost:6379", Password: ""})
	_, err := client.Ping().Result()
	if err != nil {
		return
	}
}
