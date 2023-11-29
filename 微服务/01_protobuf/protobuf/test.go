package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	text := &Test{
		Name:   "jack",
		Weight: 140,
		Height: []int32{160, 175, 180},
		Motto:  "hello",
	}

	fmt.Println(text.String())

	//proto编码
	data, err := proto.Marshal(text)
	if err != nil {
		fmt.Printf("编码失败")
	}
	//编码后的打印
	fmt.Println(data)

	newtext := &Test{}
	//解码
	err = proto.Unmarshal(data, newtext)
	if err != nil {
		fmt.Printf("解码失败")
	}
	//fmt.Println(1111)
	fmt.Println(newtext)

	fmt.Println(newtext.Motto)
}
