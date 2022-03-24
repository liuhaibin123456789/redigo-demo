package model

//Post 服务端发布帖子的结构体
type Post struct {
	UserName string `json:"user_name" form:"user_name"`
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	Number   int64  `json:"number" form:"number"`
	Time     string `json:"time" form:"time"`
}

var Channel = "c2"
