package service

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"redis/second/model"
	"redis/second/tool"
)

func CreateBroadcast(post model.Post) error {
	err := tool.LinkRedis()
	if err != nil {
		return err
	}
	defer func(conn redis.Conn) {
		err := tool.Close(conn)
		if err != nil {
			return
		}
	}(tool.RDS)
	//序列化json
	bytes, err := json.Marshal(&post)
	if err != nil {
		return err
	}
	//获取频道，选择想那些频道进行发布订阅
	_, err = redis.Bool(tool.RDS.Do("publish", model.Channel, bytes))
	return err
}

func GetBroadcast() (model.Post, error) {
	err := tool.LinkRedis()
	if err != nil {
		return model.Post{}, err
	}
	defer func(conn redis.Conn) {
		err := tool.Close(conn)
		if err != nil {
			return
		}
	}(tool.RDS)
	var post = model.Post{}
	//使用包装好发布订阅对象
	pubSubConn := redis.PubSubConn{Conn: tool.RDS}
	err = pubSubConn.Subscribe(model.Channel)
	if err != nil {
		return post, err
	}
	for true {
		switch v := pubSubConn.Receive().(type) {
		case redis.Message:
			fmt.Println("2:", v)
			err = json.Unmarshal(v.Data, &post)
			if err != nil {
				return model.Post{}, err
			}
			return post, nil
		case redis.Subscription:
			fmt.Println("1:", v)
		case error:
			fmt.Println("3:", v)
			return model.Post{}, v
		}
	}
	return post, nil
}
