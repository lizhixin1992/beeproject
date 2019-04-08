package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int
	Name    string
	ProfileId int
}

type Profile struct {
	Id   int
	Age  int16
	//userId int
}

type Post struct {
	Id    int
	Title string
	//userId  int
	//tagId  int
}

type Tag struct {
	Id    int
	Name  string
	//postId int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
}
