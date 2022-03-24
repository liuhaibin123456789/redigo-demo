package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"redis/second/api"
	_ "redis/second/docs"
)

// @title           发布与订阅
// @version         1.0
// @host            localhost:8081
// @description     基于redigo库的发布与订阅
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //注册swagger路由
	router.POST("/broadcast", api.CreateBroadcast)                            //服务端来发布帖子至频道
	router.GET("/broadcast", api.GetBroadcast)                                //客户订阅的10s内要发布订阅，否则redis连接断开
	err := router.Run(":8081")
	if err != nil {
		return
	}
}
