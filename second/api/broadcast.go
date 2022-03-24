package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redis/second/model"
	"redis/second/service"
)

// CreateBroadcast
// @summary  服务端来发布帖子至频道
// @accept mpfd
// @produce  json
// @Param    user_name     formData  string  true  "用户名"
// @Param    title     formData  string  true  "帖子标题"
// @Param    content     formData  string  true  "帖子内容"
// @Param    number     formData  string  true  "帖子关注数量"
// @Param    time     formData  string  true  "帖子发布时间"
// @router   /broadcast [post]
func CreateBroadcast(c *gin.Context) {
	//获取帖子数据
	var post model.Post
	err := c.ShouldBind(&post)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = service.CreateBroadcast(post)
	c.JSON(200, gin.H{
		"err": err,
	})

}

// GetBroadcast
// @summary  客户订阅的10s内要发布订阅的帖子
// @produce  json
// @success  200       {object}  model.Post  "成功"
// @failure  200       {object}  string  "请求错误"
// @router   /broadcast [get]
func GetBroadcast(c *gin.Context) {
	broadcast, err := service.GetBroadcast()
	if err != nil {
		c.JSON(200, gin.H{
			"err":       err,
			"subscribe": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"subscribe": broadcast,
	})
}
