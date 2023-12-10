package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"protobuf/proto/userService"
)

func main() {
	//fmt.Println("2222")
	u := &userService.Userinfo{
		Username: "张三",
		Age:      20,
		Hobby:    []string{"吃饭", "睡觉", "写代码"},
	}
	fmt.Println(u.GetType())
	fmt.Println(u.GetHobby())
	data, _ := proto.Marshal(u)
	fmt.Println(data)
	var user userService.Userinfo
	proto.Unmarshal(data, &user)
	fmt.Println(user)
	fmt.Printf("%#v\n", user)
}
